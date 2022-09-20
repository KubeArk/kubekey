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

var RookCommon = template.Must(template.New("common.yaml").Parse(
	dedent.Dedent(`
apiVersion: v1
kind: Namespace
metadata:
  name: rook-ceph # namespace:cluster
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-nodeplugin
rules:
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["namespaces"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list"]
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-external-provisioner-runner
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete", "update", "patch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments/status"]
    verbs: ["patch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims/status"]
    verbs: ["update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents"]
    verbs: ["create", "get", "list", "watch", "update", "delete", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents/status"]
    verbs: ["update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots/status"]
    verbs: ["update", "patch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: 'psp:rook'
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
rules:
  - apiGroups:
      - policy
    resources:
      - podsecuritypolicies
    resourceNames:
      - 00-rook-privileged
    verbs:
      - use
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-nodeplugin
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["namespaces"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["serviceaccounts"]
    verbs: ["get"]
  - apiGroups: [""]
    resources: ["serviceaccounts/token"]
    verbs: ["create"]
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-external-provisioner-runner
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete", "update", "patch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments/status"]
    verbs: ["patch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents"]
    verbs: ["create", "get", "list", "watch", "update", "delete", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents/status"]
    verbs: ["update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots/status"]
    verbs: ["update", "patch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims/status"]
    verbs: ["update", "patch"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get"]
  - apiGroups: ["replication.storage.openshift.io"]
    resources: ["volumereplications", "volumereplicationclasses"]
    verbs: ["create", "delete", "get", "list", "patch", "update", "watch"]
  - apiGroups: ["replication.storage.openshift.io"]
    resources: ["volumereplications/finalizers"]
    verbs: ["update"]
  - apiGroups: ["replication.storage.openshift.io"]
    resources: ["volumereplications/status"]
    verbs: ["get", "patch", "update"]
  - apiGroups: ["replication.storage.openshift.io"]
    resources: ["volumereplicationclasses/status"]
    verbs: ["get"]
  - apiGroups: [""]
    resources: ["serviceaccounts"]
    verbs: ["get"]
  - apiGroups: [""]
    resources: ["serviceaccounts/token"]
    verbs: ["create"]
---
# The cluster role for managing all the cluster-specific resources in a namespace
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: rook-ceph-cluster-mgmt
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
rules:
  - apiGroups:
      - ""
      - apps
      - extensions
    resources:
      - secrets
      - pods
      - pods/log
      - services
      - configmaps
      - deployments
      - daemonsets
    verbs:
      - get
      - list
      - watch
      - patch
      - create
      - update
      - delete
---
# The cluster role for managing the Rook CRDs
apiVersion: rbac.authorization.k8s.io/v1
# Rook watches for its CRDs in all namespaces, so this should be a cluster-scoped role unless the
# operator config ROOK_CURRENT_NAMESPACE_ONLY=true.
kind: ClusterRole
metadata:
  name: rook-ceph-global
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
rules:
  - apiGroups:
      - ""
    resources:
      # Pod access is needed for fencing
      - pods
      # Node access is needed for determining nodes where mons should run
      - nodes
      - nodes/proxy
      - services
      # Rook watches secrets which it uses to configure access to external resources.
      # e.g., external Ceph cluster; TLS certificates for the admission controller or object store
      - secrets
      # Rook watches for changes to the rook-operator-config configmap
      - configmaps
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - ""
    resources:
      # Rook creates events for its custom resources
      - events
      # Rook creates PVs and PVCs for OSDs managed by the Rook provisioner
      - persistentvolumes
      - persistentvolumeclaims
      # Rook creates endpoints for mgr and object store access
      - endpoints
    verbs:
      - get
      - list
      - watch
      - patch
      - create
      - update
      - delete
  - apiGroups:
      - storage.k8s.io
    resources:
      - storageclasses
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - batch
    resources:
      - jobs
      - cronjobs
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - delete
  # The Rook operator must be able to watch all ceph.rook.io resources to reconcile them.
  - apiGroups: ["ceph.rook.io"]
    resources:
      - cephclients
      - cephclusters
      - cephblockpools
      - cephfilesystems
      - cephnfses
      - cephobjectstores
      - cephobjectstoreusers
      - cephobjectrealms
      - cephobjectzonegroups
      - cephobjectzones
      - cephbuckettopics
      - cephbucketnotifications
      - cephrbdmirrors
      - cephfilesystemmirrors
      - cephfilesystemsubvolumegroups
      - cephblockpoolradosnamespaces
    verbs:
      - get
      - list
      - watch
      # Ideally the update permission is not required, but Rook needs it to add finalizers to resources.
      - update
  # Rook must have update access to status subresources for its custom resources.
  - apiGroups: ["ceph.rook.io"]
    resources:
      - cephclients/status
      - cephclusters/status
      - cephblockpools/status
      - cephfilesystems/status
      - cephnfses/status
      - cephobjectstores/status
      - cephobjectstoreusers/status
      - cephobjectrealms/status
      - cephobjectzonegroups/status
      - cephobjectzones/status
      - cephbuckettopics/status
      - cephbucketnotifications/status
      - cephrbdmirrors/status
      - cephfilesystemmirrors/status
      - cephfilesystemsubvolumegroups/status
      - cephblockpoolradosnamespaces/status
    verbs: ["update"]
  # The "*/finalizers" permission may need to be strictly given for K8s clusters where
  # OwnerReferencesPermissionEnforcement is enabled so that Rook can set blockOwnerDeletion on
  # resources owned by Rook CRs (e.g., a Secret owned by an OSD Deployment). See more:
  # https://kubernetes.io/docs/reference/access-authn-authz/_print/#ownerreferencespermissionenforcement
  - apiGroups: ["ceph.rook.io"]
    resources:
      - cephclients/finalizers
      - cephclusters/finalizers
      - cephblockpools/finalizers
      - cephfilesystems/finalizers
      - cephnfses/finalizers
      - cephobjectstores/finalizers
      - cephobjectstoreusers/finalizers
      - cephobjectrealms/finalizers
      - cephobjectzonegroups/finalizers
      - cephobjectzones/finalizers
      - cephbuckettopics/finalizers
      - cephbucketnotifications/finalizers
      - cephrbdmirrors/finalizers
      - cephfilesystemmirrors/finalizers
      - cephfilesystemsubvolumegroups/finalizers
      - cephblockpoolradosnamespaces/finalizers
    verbs: ["update"]
  - apiGroups:
      - policy
      - apps
      - extensions
    resources:
      # This is for the clusterdisruption controller
      - poddisruptionbudgets
      # This is for both clusterdisruption and nodedrain controllers
      - deployments
      - replicasets
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - delete
      - deletecollection
  - apiGroups:
      - healthchecking.openshift.io
    resources:
      - machinedisruptionbudgets
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - delete
  - apiGroups:
      - machine.openshift.io
    resources:
      - machines
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - delete
  - apiGroups:
      - storage.k8s.io
    resources:
      - csidrivers
    verbs:
      - create
      - delete
      - get
      - update
  - apiGroups:
      - k8s.cni.cncf.io
    resources:
      - network-attachment-definitions
    verbs:
      - get
---
# Aspects of ceph-mgr that require cluster-wide access
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-mgr-cluster
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
rules:
  - apiGroups:
      - ""
    resources:
      - configmaps
      - nodes
      - nodes/proxy
      - persistentvolumes
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - ""
    resources:
      - events
    verbs:
      - create
      - patch
      - list
      - get
      - watch
  - apiGroups:
      - storage.k8s.io
    resources:
      - storageclasses
    verbs:
      - get
      - list
      - watch
---
# Aspects of ceph-mgr that require access to the system namespace
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-mgr-system
rules:
  - apiGroups:
      - ""
    resources:
      - configmaps
    verbs:
      - get
      - list
      - watch
---
# Used for provisioning ObjectBuckets (OBs) in response to ObjectBucketClaims (OBCs).
# Note: Rook runs a copy of the lib-bucket-provisioner's OBC controller.
# OBCs can be created in any Kubernetes namespace, so this must be a cluster-scoped role.
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-object-bucket
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
rules:
  - apiGroups: [""]
    resources: ["secrets", "configmaps"]
    verbs:
      # OBC controller creates secrets and configmaps containing information for users about how to
      # connect to object buckets. It deletes them when an OBC is deleted.
      - get
      - create
      - update
      - delete
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs:
      # OBC controller gets parameters from the OBC's storageclass
      # Rook gets additional parameters from the OBC's storageclass
      - get
  - apiGroups: ["objectbucket.io"]
    resources: ["objectbucketclaims"]
    verbs:
      # OBC controller needs to list/watch OBCs and get latest version of a reconciled OBC
      - list
      - watch
      - get
      # Ideally, update should not be needed, but the OBC controller updates the OBC with bucket
      # information outside of the status subresource
      - update
      # OBC controller does not delete OBCs; users do this
  - apiGroups: ["objectbucket.io"]
    resources: ["objectbuckets"]
    verbs:
      # OBC controller needs to list/watch OBs and get latest version of a reconciled OB
      - list
      - watch
      - get
      # OBC controller creates an OB when an OBC's bucket has been provisioned by Ceph, updates them
      # when an OBC is updated, and deletes them when the OBC is de-provisioned.
      - create
      - update
      - delete
  - apiGroups: ["objectbucket.io"]
    resources: ["objectbucketclaims/status", "objectbuckets/status"]
    verbs:
      # OBC controller updates OBC and OB statuses
      - update
  - apiGroups: ["objectbucket.io"]
    # This does not strictly allow the OBC/OB controllers to update finalizers. That is handled by
    # the direct "update" permissions above. Instead, this allows Rook's controller to create
    # resources which are owned by OBs/OBCs and where blockOwnerDeletion is set.
    resources: ["objectbucketclaims/finalizers", "objectbuckets/finalizers"]
    verbs:
      - update
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-osd
rules:
  - apiGroups:
      - ""
    resources:
      - nodes
    verbs:
      - get
      - list
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-system
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
rules:
  # Most resources are represented by a string representation of their name, such as "pods", just as it appears in the URL for the relevant API endpoint.
  # However, some Kubernetes APIs involve a "subresource", such as the logs for a pod. [...]
  # To represent this in an RBAC role, use a slash to delimit the resource and subresource.
  # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources
  - apiGroups: [""]
    resources: ["pods", "pods/log"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["pods/exec"]
    verbs: ["create"]
  - apiGroups: ["admissionregistration.k8s.io"]
    resources: ["validatingwebhookconfigurations"]
    verbs: ["create", "get", "delete", "update"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-nodeplugin
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-plugin-sa
    namespace: rook-ceph # namespace:operator
roleRef:
  kind: ClusterRole
  name: cephfs-csi-nodeplugin
  apiGroup: rbac.authorization.k8s.io
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-provisioner-role
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-provisioner-sa
    namespace: rook-ceph # namespace:operator
roleRef:
  kind: ClusterRole
  name: cephfs-external-provisioner-runner
  apiGroup: rbac.authorization.k8s.io
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-nodeplugin
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-plugin-sa
    namespace: rook-ceph # namespace:operator
roleRef:
  kind: ClusterRole
  name: rbd-csi-nodeplugin
  apiGroup: rbac.authorization.k8s.io
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-provisioner-role
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-provisioner-sa
    namespace: rook-ceph # namespace:operator
roleRef:
  kind: ClusterRole
  name: rbd-external-provisioner-runner
  apiGroup: rbac.authorization.k8s.io
---
# Grant the rook system daemons cluster-wide access to manage the Rook CRDs, PVCs, and storage classes
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-global
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-global
subjects:
  - kind: ServiceAccount
    name: rook-ceph-system
    namespace: rook-ceph # namespace:operator
---
# Allow the ceph mgr to access cluster-wide resources necessary for the mgr modules
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-mgr-cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-mgr-cluster
subjects:
  - kind: ServiceAccount
    name: rook-ceph-mgr
    namespace: rook-ceph # namespace:cluster
---
kind: ClusterRoleBinding
# Give Rook-Ceph Operator permissions to provision ObjectBuckets in response to ObjectBucketClaims.
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-object-bucket
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-object-bucket
subjects:
  - kind: ServiceAccount
    name: rook-ceph-system
    namespace: rook-ceph # namespace:operator
---
# Allow the ceph osd to access cluster-wide resources necessary for determining their topology location
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-osd
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-osd
subjects:
  - kind: ServiceAccount
    name: rook-ceph-osd
    namespace: rook-ceph # namespace:cluster
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-system
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-system
subjects:
  - kind: ServiceAccount
    name: rook-ceph-system
    namespace: rook-ceph # namespace:operator
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-ceph-system-psp
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-ceph-system
    namespace: rook-ceph # namespace:operator
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-csi-cephfs-plugin-sa-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-plugin-sa
    namespace: rook-ceph # namespace:operator
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-csi-cephfs-provisioner-sa-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-provisioner-sa
    namespace: rook-ceph # namespace:operator
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-csi-rbd-plugin-sa-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-plugin-sa
    namespace: rook-ceph # namespace:operator
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-csi-rbd-provisioner-sa-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-provisioner-sa
    namespace: rook-ceph # namespace:operator
---
# We expect most Kubernetes teams to follow the Kubernetes docs and have these PSPs.
# * privileged (for kube-system namespace)
# * restricted (for all logged in users)
#
# PSPs are applied based on the first match alphabetically. rook-ceph-operator comes after
# restricted alphabetically, so we name this 00-rook-privileged, so it stays somewhere
# close to the top and so rook-system gets the intended PSP. This may need to be renamed in
# environments with other 00-prefixed PSPs.
#
# More on PSP ordering: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#policy-order
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: 00-rook-privileged
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: 'runtime/default'
    seccomp.security.alpha.kubernetes.io/defaultProfileName: 'runtime/default'
spec:
  privileged: true
  allowedCapabilities:
    # required by CSI
    - SYS_ADMIN
    - MKNOD
  fsGroup:
    rule: RunAsAny
  # runAsUser, supplementalGroups - Rook needs to run some pods as root
  # Ceph pods could be run as the Ceph user, but that user isn't always known ahead of time
  runAsUser:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  # seLinux - seLinux context is unknown ahead of time; set if this is well-known
  seLinux:
    rule: RunAsAny
  volumes:
    # recommended minimum set
    - configMap
    - downwardAPI
    - emptyDir
    - persistentVolumeClaim
    - secret
    - projected
    # required for Rook
    - hostPath
  # allowedHostPaths can be set to Rook's known host volume mount points when they are fully-known
  # allowedHostPaths:
  #   - pathPrefix: "/run/udev"  # for OSD prep
  #     readOnly: false
  #   - pathPrefix: "/dev"  # for OSD prep
  #     readOnly: false
  #   - pathPrefix: "/var/lib/rook"  # or whatever the dataDirHostPath value is set to
  #     readOnly: false
  # Ceph requires host IPC for setting up encrypted devices
  hostIPC: true
  # Ceph OSDs need to share the same PID namespace
  hostPID: true
  # hostNetwork can be set to 'false' if host networking isn't used
  hostNetwork: true
  hostPorts:
    # Ceph messenger protocol v1
    - min: 6789
      max: 6790 # <- support old default port
    # Ceph messenger protocol v2
    - min: 3300
      max: 3300
    # Ceph RADOS ports for OSDs, MDSes
    - min: 6800
      max: 7300
    # # Ceph dashboard port HTTP (not recommended)
    # - min: 7000
    #   max: 7000
    # Ceph dashboard port HTTPS
    - min: 8443
      max: 8443
    # Ceph mgr Prometheus Metrics
    - min: 9283
      max: 9283
    # port for CSIAddons
    - min: 9070
      max: 9070
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-external-provisioner-cfg
  namespace: rook-ceph # namespace:operator
rules:
  - apiGroups: [""]
    resources: ["endpoints"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list", "create", "delete"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-nodeplugin
  namespace: rook-ceph # namespace:operator
rules:
  - apiGroups: ["csiaddons.openshift.io"]
    resources: ["csiaddonsnodes"]
    verbs: ["create"]
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-external-provisioner-cfg
  namespace: rook-ceph # namespace:operator
rules:
  - apiGroups: [""]
    resources: ["endpoints"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list", "watch", "create", "delete", "update"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
  - apiGroups: ["csiaddons.openshift.io"]
    resources: ["csiaddonsnodes"]
    verbs: ["create"]
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-cmd-reporter
  namespace: rook-ceph # namespace:cluster
rules:
  - apiGroups:
      - ""
    resources:
      - pods
      - configmaps
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - delete
---
# Aspects of ceph-mgr that operate within the cluster's namespace
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-mgr
  namespace: rook-ceph # namespace:cluster
rules:
  - apiGroups:
      - ""
    resources:
      - pods
      - services
      - pods/log
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - delete
  - apiGroups:
      - batch
    resources:
      - jobs
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - delete
  - apiGroups:
      - ceph.rook.io
    resources:
      - "*"
    verbs:
      - "*"
  - apiGroups:
      - apps
    resources:
      - deployments/scale
      - deployments
    verbs:
      - patch
      - delete
  - apiGroups:
      - ''
    resources:
      - persistentvolumeclaims
    verbs:
      - delete
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-osd
  namespace: rook-ceph # namespace:cluster
rules:
  # this is needed for rook's "key-management" CLI to fetch the vault token from the secret when
  # validating the connection details
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list", "watch", "create", "update", "delete"]
  - apiGroups: ["ceph.rook.io"]
    resources: ["cephclusters", "cephclusters/finalizers"]
    verbs: ["get", "list", "create", "update", "delete"]
---
# Aspects of ceph osd purge job that require access to the cluster namespace
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-purge-osd
  namespace: rook-ceph # namespace:cluster
rules:
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get"]
  - apiGroups: ["apps"]
    resources: ["deployments"]
    verbs: ["get", "delete"]
  - apiGroups: ["batch"]
    resources: ["jobs"]
    verbs: ["get", "list", "delete"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "update", "delete", "list"]
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-rgw
  namespace: rook-ceph # namespace:cluster
rules:
  # Placeholder role so the rgw service account will
  # be generated in the csv. Remove this role and role binding
  # when fixing https://github.com/rook/rook/issues/10141.
  - apiGroups:
      - ""
    resources:
      - configmaps
    verbs:
      - get
---
# Allow the operator to manage resources in its own namespace
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: rook-ceph-system
  namespace: rook-ceph # namespace:operator
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
rules:
  - apiGroups:
      - ""
    resources:
      - pods
      - configmaps
      - services
    verbs:
      - get
      - list
      - watch
      - patch
      - create
      - update
      - delete
  - apiGroups:
      - apps
      - extensions
    resources:
      - daemonsets
      - statefulsets
      - deployments
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - delete
  - apiGroups:
      - batch
    resources:
      - cronjobs
    verbs:
      - delete
  - apiGroups:
      - cert-manager.io
    resources:
      - certificates
      - issuers
    verbs:
      - get
      - create
      - delete
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-provisioner-role-cfg
  namespace: rook-ceph # namespace:operator
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-provisioner-sa
    namespace: rook-ceph # namespace:operator
roleRef:
  kind: Role
  name: cephfs-external-provisioner-cfg
  apiGroup: rbac.authorization.k8s.io
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-nodeplugin-role-cfg
  namespace: rook-ceph # namespace:operator
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-plugin-sa
    namespace: rook-ceph # namespace:operator
roleRef:
  kind: Role
  name: rbd-csi-nodeplugin
  apiGroup: rbac.authorization.k8s.io
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-provisioner-role-cfg
  namespace: rook-ceph # namespace:operator
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-provisioner-sa
    namespace: rook-ceph # namespace:operator
roleRef:
  kind: Role
  name: rbd-external-provisioner-cfg
  apiGroup: rbac.authorization.k8s.io
---
# Allow the operator to create resources in this cluster's namespace
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-cluster-mgmt
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-cluster-mgmt
subjects:
  - kind: ServiceAccount
    name: rook-ceph-system
    namespace: rook-ceph # namespace:operator
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-cmd-reporter
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-cmd-reporter
subjects:
  - kind: ServiceAccount
    name: rook-ceph-cmd-reporter
    namespace: rook-ceph # namespace:cluster
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-cmd-reporter-psp
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
  - kind: ServiceAccount
    name: rook-ceph-cmd-reporter
    namespace: rook-ceph # namespace:cluster
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-default-psp
  namespace: rook-ceph # namespace:cluster
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
  - kind: ServiceAccount
    name: default
    namespace: rook-ceph # namespace:cluster
---
# Allow the ceph mgr to access resources scoped to the CephCluster namespace necessary for mgr modules
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-mgr
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-mgr
subjects:
  - kind: ServiceAccount
    name: rook-ceph-mgr
    namespace: rook-ceph # namespace:cluster
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-mgr-psp
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
  - kind: ServiceAccount
    name: rook-ceph-mgr
    namespace: rook-ceph # namespace:cluster
---
# Allow the ceph mgr to access resources in the Rook operator namespace necessary for mgr modules
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-mgr-system
  namespace: rook-ceph # namespace:operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-mgr-system
subjects:
  - kind: ServiceAccount
    name: rook-ceph-mgr
    namespace: rook-ceph # namespace:cluster
---
# Allow the osd pods in this namespace to work with configmaps
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-osd
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-osd
subjects:
  - kind: ServiceAccount
    name: rook-ceph-osd
    namespace: rook-ceph # namespace:cluster
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-osd-psp
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
  - kind: ServiceAccount
    name: rook-ceph-osd
    namespace: rook-ceph # namespace:cluster
---
# Allow the osd purge job to run in this namespace
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-purge-osd
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-purge-osd
subjects:
  - kind: ServiceAccount
    name: rook-ceph-purge-osd
    namespace: rook-ceph # namespace:cluster
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-purge-osd-psp
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
  - kind: ServiceAccount
    name: rook-ceph-purge-osd
    namespace: rook-ceph # namespace:cluster
---
# Allow the rgw pods in this namespace to work with configmaps
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-rgw
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-rgw
subjects:
  - kind: ServiceAccount
    name: rook-ceph-rgw
    namespace: rook-ceph # namespace:cluster
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-rgw-psp
  namespace: rook-ceph # namespace:cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
  - kind: ServiceAccount
    name: rook-ceph-rgw
    namespace: rook-ceph # namespace:cluster
---
# Grant the operator, agent, and discovery agents access to resources in the rook-ceph-system namespace
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rook-ceph-system
  namespace: rook-ceph # namespace:operator
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-system
subjects:
  - kind: ServiceAccount
    name: rook-ceph-system
    namespace: rook-ceph # namespace:operator
---
# Service account for the job that reports the Ceph version in an image
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-cmd-reporter
  namespace: rook-ceph # namespace:cluster
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for Ceph mgrs
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-mgr
  namespace: rook-ceph # namespace:cluster
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for Ceph OSDs
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-osd
  namespace: rook-ceph # namespace:cluster
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for job that purges OSDs from a Rook-Ceph cluster
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-purge-osd
  namespace: rook-ceph # namespace:cluster
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for RGW server
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-rgw
  namespace: rook-ceph # namespace:cluster
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for the Rook-Ceph operator
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-system
  namespace: rook-ceph # namespace:operator
  labels:
    operator: rook
    storage-backend: ceph
    app.kubernetes.io/part-of: rook-ceph-operator
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for the CephFS CSI driver
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-csi-cephfs-plugin-sa
  namespace: rook-ceph # namespace:operator
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for the CephFS CSI provisioner
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-csi-cephfs-provisioner-sa
  namespace: rook-ceph # namespace:operator
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for the RBD CSI driver
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-csi-rbd-plugin-sa
  namespace: rook-ceph # namespace:operator
# imagePullSecrets:
#   - name: my-registry-secret
---
# Service account for the RBD CSI provisioner
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-csi-rbd-provisioner-sa
  namespace: rook-ceph # namespace:operator
# imagePullSecrets:
#   - name: my-registry-secret
`)))