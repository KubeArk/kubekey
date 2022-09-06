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

var KubearkAppManifest = template.Must(template.New("kubeark.yaml").Parse(
	dedent.Dedent(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kubeark-web-frontend
  namespace: kubeark
  labels:
    app: kubeark-web-frontend
    namespace: kubeark
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kubeark-web-frontend
      namespace: kubeark
  template:
    metadata:
      labels:
        app: kubeark-web-frontend
        namespace: kubeark
    spec:
      containers:
        - name: kubeark-frontend
          image: "kubeark/kubeark-frontend:latest"
          ports:
            - containerPort: 3000
              protocol: TCP
          env:
            - name: BACKEND_API
              value: "http://kubeark-web-app-service:80/api/v1/"
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          imagePullPolicy: Always
      restartPolicy: Always
      terminationGracePeriodSeconds: 0
      dnsPolicy: ClusterFirst
      securityContext: {}
      imagePullSecrets:
        - name: kubeark-docker-hub
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 1
  minReadySeconds: 5
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
---
apiVersion: v1
kind: Service
metadata:
  name: kubeark-web-frontend-service
  namespace: kubeark
  labels:
    app: kubeark-web-frontend
    namespace: kubeark
spec:
  ports:
    - name: http
      protocol: TCP
      port: 80
      targetPort: 3000
  selector:
    app: kubeark-web-frontend
    namespace: kubeark
  type: LoadBalancer
  sessionAffinity: None
  externalTrafficPolicy: Cluster
  ipFamilies:
    - IPv4
  ipFamilyPolicy: SingleStack
  internalTrafficPolicy: Cluster
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: kubeark-web-scheduler
  namespace: kubeark
  labels:
    app: kubeark-web-scheduler
    namespace: kubeark
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kubeark-web-scheduler
      namespace: kubeark
  template:
    metadata:
      labels:
        app: kubeark-web-scheduler
        namespace: kubeark
    spec:
      containers:
        - name: kubeark
          image: 'kubeark/kubeark:ee22560'
          command: ["python"]
          args: ["scheduler.py"]
          envFrom:
            - configMapRef:
                name: kubeark-config
            - secretRef:
                name: postgresscluster-pguser-kube-db
          env:
          - name: CONTAINER_ROLE
            value: scheduler
          ports:
            - containerPort: 8000
              protocol: TCP
          resources: {}
          imagePullPolicy: Always
          volumeMounts:
            - name: kubeark-web-app-volume
              mountPath: kubeark-backend/app/storage/
      volumes:
        - name: kubeark-web-app-volume
          persistentVolumeClaim:
            claimName: kubeark-storage-pvc
      restartPolicy: Always
      terminationGracePeriodSeconds: 0
      securityContext: {}
      imagePullSecrets:
        - name: kubeark-docker-hub
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 1
  minReadySeconds: 5
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: kubeark-web-app
  namespace: kubeark
  labels:
    app: kubeark-web-app
    namespace: kubeark
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kubeark-web-app
      namespace: kubeark
  template:
    metadata:
      labels:
        app: kubeark-web-app
        namespace: kubeark
    spec:
      containers:
        - name: kubeark
          image: 'kubeark/kubeark:latest'
          envFrom:
            - configMapRef:
                name: kubeark-config
          env:
          - name: CONTAINER_ROLE
            value: app
          ports:
            - containerPort: 8000
              protocol: TCP
          resources: {}
          imagePullPolicy: Always
          volumeMounts:
            - name: kubeark-web-app-volume
              mountPath: kubeark-backend/app/storage/
      volumes:
        - name: kubeark-web-app-volume
          persistentVolumeClaim:
            claimName: kubeark-storage-pvc
      restartPolicy: Always
      terminationGracePeriodSeconds: 0
      securityContext: {}
      imagePullSecrets:
        - name: kubeark-docker-hub
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 1
  minReadySeconds: 5
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
---
kind: Service
apiVersion: v1
metadata:
  name: kubeark-web-app-service
  namespace: kubeark
  labels:
    app: kubeark-web-app
    namespace: kubeark
spec:
  selector:
    app: kubeark-web-app
    namespace: kubeark
  ports:
    - name: http
      port: 80
      targetPort: 8000
  type: LoadBalancer
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: kubeark-web-worker
  namespace: kubeark
  labels:
    app: kubeark-web-worker
    namespace: kubeark
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kubeark-web-worker
      namespace: kubeark
  template:
    metadata:
      labels:
        app: kubeark-web-worker
        namespace: kubeark
    spec:
      containers:
        - name: kubeark
          image: 'kubeark/kubeark:latest'
          command: ["python"]
          args: ["worker.py"]
          envFrom:
            - configMapRef:
                name: kubeark-config
          env:
          - name: CONTAINER_ROLE
            value: worker
          ports:
            - containerPort: 8000
              protocol: TCP
          resources: {}
          imagePullPolicy: Always
          volumeMounts:
            - name: kubeark-web-app-volume
              mountPath: kubeark-backend/app/storage/
      volumes:
        - name: kubeark-web-app-volume
          persistentVolumeClaim:
            claimName: kubeark-storage-pvc
      restartPolicy: Always
      terminationGracePeriodSeconds: 0
      securityContext: {}
      imagePullSecrets:
        - name: kubeark-docker-hub
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
      maxSurge: 1
  minReadySeconds: 5
  revisionHistoryLimit: 10
  progressDeadlineSeconds: 600
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis
  namespace: redis
  labels:
    app: redis
    tier: kubeark-backend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
      tier: kubeark-backend
  template:
    metadata:
      labels:
        app: redis
        tier: kubeark-backend
    spec:
      containers:
      - name: redis
        image: "docker.io/redis:6.0.5"
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
        ports:
        - containerPort: 6379
---
apiVersion: v1
kind: Service
metadata:
  name: kubeark-redis-service
  namespace: redis
  labels:
    app: redis
    tier: kubeark-backend
spec:
  ports:
    - name: http
      protocol: TCP
      port: 6379
      targetPort: 6379
  selector:
    app: redis
    tier: kubeark-backend
  type: ClusterIP
  sessionAffinity: None
  ipFamilies:
    - IPv4
  ipFamilyPolicy: SingleStack
`)))
