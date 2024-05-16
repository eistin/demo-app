# Demo Microservices Counter App

This repository contains a demo microservices application implementing a simple counter app. The application consists of separate backend and frontend services, along with supporting components like a database and Nginx server.

## Overview
The counter app is designed to demonstrate the interaction between multiple microservices in a containerized environment. The backend service, implemented in Go, provides APIs to increment and retrieve the counter value. The frontend, written in PHP, interacts with the backend APIs to display the current counter value and allows users to increment it.

## Project Structure
```
.
├── README.md                  # This file
├── docker-compose.yml         # Docker Compose configuration for local deployment
├── manifests                  # Kubernetes manifests for deployment with ArgoCD
│   └── k8s
│       ├── backend.yaml       # Backend service deployment YAML
│       ├── frontend.yaml      # Frontend service deployment YAML
└── micro-services             # Source code and configurations for microservices
    ├── backend
    │   ├── Dockerfile.backend  # Dockerfile for backend service
    │   ├── go.mod              # Go module file
    │   ├── go.sum              # Go dependencies checksum file
    │   └── main.go             # Main backend service code
    ├── db
    │   └── schema.sql          # SQL schema for database
    ├── frontend
    │   ├── Dockerfile.frontend # Dockerfile for frontend service
    │   ├── index.php           # Main PHP file for frontend
    │   └── styles.css          # CSS styles for frontend
    └── nginx
        └── nginx.conf          # Nginx configuration file
```

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/engine/install/)
  
### Local Deployment

1. Clone this repository:
```
git clone https://github.com/eistin/demo-app.git
```
2. Navigate to the project directory:
```
cd demo-app/
```
3. Run the following command to start the application locally:
```
docker-compose up --build
```
4. Access the application in your browser at `http://localhost:8000`.

### Deployment with ArgoCD

You can use my other repository [demo-ansible-k8s](https://github.com/eistin/demo-ansible-k8s) to deploy argocd and the demo-app into your local cluster.