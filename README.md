# AoC 2021

[Advent of Code 2021](https://adventofcode.com/2021) but each solution is a horribly-written microservice
and the entire thing is deployed with Kubernetes for some reason.

## Development

Install Docker, kubectl, minikube.

Create deployment:
```
minikube start

# Enable ingress
minikube addons enable ingress

# Build images
# Use minikube image repository (Windows-specific)
minikube -p minikube docker-env | Invoke-Expression
./scripts/build_images.sh

# Create deployments
./scripts/deploy.sh
```

Tunnel to ingress at `localhost`:
```
minikube tunnel
```