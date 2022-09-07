/*
 Copyright 2022 The KubeSphere Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package templates

import (
	"text/template"

	"github.com/lithammer/dedent"
)

var KubearkConfigs = template.Must(template.New("kubeark-configs.yaml").Parse(
	dedent.Dedent(`
apiVersion: v1
kind: Namespace
metadata:
  name: kubeark
---
apiVersion: v1
kind: Namespace
metadata:
  name: redis
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: kubeark-config
  namespace: kubeark
  labels:
    name: kubeark-config
data:
  FLASK_ENV: "dev"
  FLASK_APP: "run"
  PYTHONUNBUFFERED: "1"
  LOG_TYPE: "stream"
  UPLOAD_FOLDER: "/app/storage/"
  PUBLIC_CHARTS_FOLDER: "public/charts"
  PRIVATE_CHARTS_FOLDER: "private/charts"
  JWT_ACCESS_TOKEN_EXPIRES_MINUTES: "20"
  JWT_REFRESH_TOKEN_EXPIRES_MINUTES: "25"
  SQLALCHEMY_POOL_RECYCLE: "3600"
  SQLALCHEMY_POOL_TIMEOUT: "100"
  SQLALCHEMY_POOL_SIZE: "200"
  SQLALCHEMY_MAX_OVERFLOW: "200"
  MAIL_SERVER: ""
  MAIL_PASSWORD: ""
  MAIL_USE_SSL: "True"
  MAIL_USE_TLS: "False"
  MAIL_USERNAME: ""
  REPOSITORY_UPDATE: "True"
  DEPLOYMENT_MAIL_NOTIFICATION: "False"
  ALLOWED_EXTENSIONS: "tgz"
  PROMETHEUS: "prometheus-service.monitoring.svc:8080"
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: internal-kubectl
  namespace: postgres-operator
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  # "namespace" omitted since ClusterRoles are not namespaced
  name: postgres-secret-reader
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "watch", "list", "delete", "create", "apply", "update", "patch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: postgressecretslusterbinding
subjects:
  - kind: ServiceAccount
    name: internal-kubectl
    namespace: postgres-operator
roleRef:
  kind: ClusterRole
  name: postgres-secret-reader
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: batch/v1
kind: Job
metadata:
  name: copy-postgres-secret
  namespace: postgres-operator
spec:
  template:
    metadata:
      name: copy-postgres-secret
    spec:
      serviceAccountName: internal-kubectl
      containers:
      - name: kubectl
        image: bitnami/kubectl
        command:
         - "bin/bash"
         - "-c"
         #TODO change demo namespace below - make it dynamic with helm charts
         - "kubectl get secret postgresscluster-pguser-kube-db -n postgres-operator -o json | jq 'del(.metadata[\"namespace\",\"creationTimestamp\",\"resourceVersion\",\"selfLink\",\"uid\",\"ownerReferences\"])' | kubectl apply -f -"
      restartPolicy: Never
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: kubeark-webapp-ingress
  namespace: kubeark
  annotations:
    kubernetes.io/ingress.allow-http: "true"
    nginx.ingress.kubernetes.io/force-ssl-redirect: "false"
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/use-regex: "true"
    nginx.ingress.kubernetes.io/http2-push-preload: "true"
    nginx.ingress.kubernetes.io/service-upstream: "true"
    nginx.ingress.kubernetes.io/proxy-connect-timeout: 60s
    nginx.ingress.kubernetes.io/proxy-send-timeout: 60s
    nginx.ingress.kubernetes.io/proxy-read-timeout:  60s
    nginx.ingress.kubernetes.io/connection-proxy-header: "keep-alive"
    nginx.ingress.kubernetes.io/configuration-snippet: "
	  nginx.ingress.kubernetes.io/rewrite-target: /$1
      keepalive_timeout 60s;
      send_timeout 60s;"
spec:
  tls:
    - hosts:
        - {{ .IngressHost }}
      secretName: modex-main-tls
  rules:
    - host: {{ .IngressHost }}
      http:
        paths:
        - path: /
          pathType: Prefix
          backend:
            service:
              name: kubeark-web-app-service
              port:
                number: 80
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: kubeark-storage-pvc
  namespace: kubeark
spec:
  storageClassName: rook-cephfs
  accessModes:
  - ReadWriteMany
  resources:
    requests:
      storage: 20Gi
---
kind: Secret
apiVersion: v1
metadata:
  name: kubeark-docker-hub
  namespace: kubeark
data:
  .dockerconfigjson: >-
    eyJhdXRocyI6eyJkb2NrZXIuaW8iOnsidXNlcm5hbWUiOiJrdWJlYXJrIiwicGFzc3dvcmQiOiI5MThhZjRmNy0zOWFiLTQ1YjUtYTkyNy01ZDUzNjhmMTU4NmEifX19
type: kubernetes.io/dockerconfigjson
---
apiVersion: v1
kind: Secret
metadata:
  name: kubeark-secrets
  namespace: kubeark
type: Opaque
stringData:
    MAIL_SERVER: MAIL_SERVER_PLACEHOLDER
    MAIL_PASSWORD: MAIL_PASSWORD_PLACEHOLDER
    MAIL_USERNAME: MAIL_USERNAME_PLACEHOLDER
    REDIS_JWT_BLOCKLIST: REDIS_JWT_BLOCKLIST_PLACEHOLDER
    CELERY_BROKER_URL: CELERY_BROKER_URL_PLACEHOLDER
    CELERY_RESULT_BACKEND: CELERY_RESULT_BACKEND_PLACEHOLDER
    REDIS_POLLING_URL: REDIS_POLLING_URL_PLACEHOLDER
`)))
