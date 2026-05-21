# Go K3s Cloud Native Lab

Study project focused on learning Kubernetes and cloud-native fundamentals using a Go API running on a local K3s cluster.

## Stack

- Go
- Docker
- Docker Compose
- Kubernetes (K3s)
- Makefile

## Features

- Multi-stage Docker build
- Local private registry
- Kubernetes Deployment and Service
- Readiness and Liveness probes
- ConfigMap and Secret integration
- Automated local workflow with Makefile

## Project Structure

```text
.
├── api/
├── k8s/
│   ├── manifests/
│   └── registries.yaml
├── docker-compose.yaml
└── makefile
```

## Commands

```bash
make dev        # Start infrastructure and deploy application
make redeploy   # Rebuild and restart application
make logs       # View application logs
make pods       # List Kubernetes pods
make down       # Stop environment
```

## Concepts Practiced

- Pods
- Deployments
- Services
- Probes
- ConfigMaps
- Secrets
- Container lifecycle
- Kubernetes networking
- Cloud-native application design
