# store-admin-go-ws

## Table of Contents
- [Description](#description)
- [Installation](#installation)
- [Running the App](#running-the-app)
- [Test](#test)
- [Docker](#docker)
  - [Image Resource Usage Metrics](#image-resource-usage-metrics)
- [Kubernetes](#kubernetes)
  - [Pod Resource Usage Metrics](#pod-resource-usage-metrics)

## Description

Store's Admin Web Service example using [Nest](https://github.com/nestjs/nest) framework.

## Installation

```bash
go get -u github.com/swaggo/swag/cmd/swag
$ go init
```

## Running the app
The following commands allow you to run the application

```bash
# development
go run .
```

### Swagger API documentation
You can access the Swagger documentation at: `http://localhost:8080/swagger/index.html`

## Docker

```bash
# Build Docker image
docker build -t store-admin-go-ws:latest -f Dockerfile .

# Run Docker container (with example port mappings and environment variables)
docker run -p 3040:3040 -p 5432:5432 -e NODE_ENV=production -e DB_HOST="host.docker.internal" -e DB_PORT="5432" -e DB_USERNAME="postgres" -e DB_PASSWORD="1234" -e DB_NAME="sale-management-system" -e DB_LOGGING="true" store-admin-go-ws
```

### Image resource usage metrics

The table below shows resource usage metrics for the `store-admin-go-ws` Docker container.

| REPOSITORY           | TAG    | IMAGE ID      | CREATED    | SIZE    |
|----------------------|--------|---------------|------------|---------|
| store-admin-go-ws    | latest | ea98a671f394  | 6 minutes  | 35.4MB  |


## Kubernetes

```bash
# Start Minikube to create a local Kubernetes cluster
minikube start

# Configure the shell to use Minikube's Docker daemon
& minikube -p minikube docker-env --shell powershell | Invoke-Expression

# Build Docker image with a specific tag and Dockerfile
docker build -t store-admin-go-ws:latest -f Dockerfile .

# Apply Kubernetes configuration to create a pod
kubectl apply -f kubernetes/pod.yaml

# Port-forward to access the Kubernetes pod locally
kubectl port-forward store-admin-go-ws-pod 3040:3040
```

### Pod resource usage metrics

The table below shows resource usage metrics for the `store-admin-go-ws-pod` pod.

```bash
minikube addons enable metrics-server
kubectl top pods
```

**Note:** If you just enabled the metrics-server addon, remember to wait a couple of seconds before running the `kubectl top pods` command.


| NAME                   | CPU(cores) | MEMORY(bytes) |
|------------------------|------------|---------------|
| store-admin-go-ws-pod  | 1m         | 15Mi          |
