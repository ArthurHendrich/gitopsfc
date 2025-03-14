# itOps with Argo CD

This repository contains a simple Go application deployed to Kubernetes using GitOps principles with Argo CD.

## Repository Structure

```
.
├── .github/
│   └── workflows/
│       └── cd.yaml         # GitHub Actions workflow for CI/CD
├── k8s/
│   ├── deployment.yaml     # Kubernetes Deployment manifest
│   ├── service.yaml        # Kubernetes Service manifest
│   └── kustomization.yaml  # Kustomize configuration
├── Dockerfile              # Docker image definition
├── go.mod                  # Go module definition
├── main.go                 # Go application source code
└── README.md               # This file
```

## Kubernetes Manifests

### Deployment (`k8s/deployment.yaml`)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goserver
spec:
  selector:
    matchLabels:
      app: goserver
  template:
    metadata:
      labels:
        app: goserver
    spec:
      containers:
      - name: goserver
        image: goserver
        ports:
        - containerPort: 8080
```

This manifest defines a Deployment named `goserver` that runs pods with the `goserver` image and exposes port 8080.

### Service (`k8s/service.yaml`)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: goserver-service
spec:
  selector:
    app: goserver
  ports:
  - port: 8080
    targetPort: 8080
```

This manifest defines a Service named `goserver-service` that routes traffic to pods with the label `app: goserver` on port 8080.

### Kustomization (`k8s/kustomization.yaml`)

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
  - deployment.yaml
  - service.yaml

images: 
- name: goserver
  newName: goserver
  newTag: latest
```

This Kustomize configuration combines the Deployment and Service manifests and sets the image tag to `latest`.

## CI/CD Workflow

The GitHub Actions workflow in `.github/workflows/cd.yaml` performs the following steps:

1. Builds a Docker image from the application code
2. Pushes the image to Docker Hub
3. Updates the Kubernetes manifests with the new image tag
4. Commits and pushes the changes back to the repository

Argo CD then detects these changes and applies them to the Kubernetes cluster.

## Troubleshooting Argo CD Sync Issues

If Argo CD reports that resources are not found in the cluster, follow these steps to resolve the issue:

### 1. Check the Argo CD Application Status

```bash
kubectl get applications -n argocd -o wide
```

This command shows the sync status and health status of your Argo CD applications.

### 2. Update the Image Tag

If the image tag in `kustomization.yaml` is invalid, update it to a valid tag:

```bash
# Edit the kustomization.yaml file to use a valid tag
# For example, change from an invalid tag to 'latest'
```

### 3. Configure Automated Sync

To enable automated sync with pruning and self-healing:

```bash
# Remove any existing sync policy
kubectl -n argocd patch applications goserver -p '{"spec":{"syncPolicy":null}}' --type=merge

# Set up automated sync with pruning and self-healing
kubectl -n argocd patch applications goserver -p '{"spec":{"syncPolicy":{"automated":{"prune":true,"selfHeal":true}}}}' --type=merge
```

### 4. Verify the Resources

After syncing, verify that the resources have been created:

```bash
kubectl get deployments,services | grep goserver
```

## Additional Argo CD Commands

### Manually Trigger a Sync

If you're not using automated sync, you can manually trigger a sync:

```bash
# Using the Argo CD CLI
argocd app sync goserver

# Or using kubectl with the Argo CD API
kubectl -n argocd patch applications goserver -p '{"operation":{"sync":{}}}' --type=merge
```

### View Application Details

```bash
# Using the Argo CD CLI
argocd app get goserver

# Or using kubectl
kubectl -n argocd get applications goserver -o yaml
```

## Notes

- Make sure your Kubernetes cluster has access to pull the Docker images specified in your manifests
- Ensure that the Argo CD application is configured to watch the correct repository and path
- If using private repositories, configure the appropriate secrets in Argo CD
