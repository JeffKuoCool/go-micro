package template

// KubernetesEnv is a Kubernetes configmap manifest template used for
// environment variables in new projects.
var KubernetesEnv = `---

apiVersion: v1
kind: ConfigMap
metadata:
  name: {{.Dir}}-env
data:
  MICRO_REGISTRY: kubernetes
`

// KubernetesClusterRole is a Kubernetes cluster role manifest template
// required for the Kubernetes registry plugin to function correctly.
var KubernetesClusterRole = `---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: micro-registry
rules:
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - list
  - patch
  - watch
`

// KubernetesRoleBinding is a Kubernetes role binding manifest template
// required for the Kubernetes registry plugin to function correctly.
var KubernetesRoleBinding = `---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: micro-registry
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: micro-registry
subjects:
- kind: ServiceAccount
  name: default
  namespace: default
`

// KubernetesDeployment is a Kubernetes deployment manifest template used for
// new projects.
var KubernetesDeployment = `---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Dir}}
  labels:
    app: {{.Dir}}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.Dir}}
  template:
    metadata:
      labels:
        app: {{.Dir}}
    spec:
      containers:
      - name: {{.Dir}}
        image: {{.Dir}}:latest
        envFrom:
        - configMapRef:
            name: {{.Dir}}-env
`
