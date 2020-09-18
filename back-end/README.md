# back-end

## Table of Contents

* [Introduction](#introduction)

* [Getting Started](#getting-started)

* [Prerequisites](#prerequisites)

    * [Install Go tools](#install-go-tools)

    * [GitLab Variables](#gitlab-variables)

    * [Pipeline](#pipeline)

* [Testing](#testing)

* [Deployment](#deployment)

* [Helm](#helm)

* [Dockerfile](#dockerfile)

* [Final](#final)
    

## Introduction

It's a [Todo-Backend](https://todobackend.com/) project.
It works with in-memory database. 
It has only handler tests.

After build stage it runs as a [Docker container](https://www.docker.com/resources/what-container) deployed with [Helm](https://helm.sh/) to GKE Cluster.

## Getting Started

A [Golang](https://golang.org/) implementation using the [Echo Web Framework](https://echo.labstack.com/), courtesy of [Simar Kalra](https://github.com/theramis/todo-backend-go-echo).

[Echo](https://echo.labstack.com/) is High performance, extensible, minimalist **Go** web framework.

[Docker](https://www.docker.com/) is a tool designed to make it easier to create, deploy, and run applications by using containers. 

[Helm](https://helm.sh/) is the package manager for Kubernetes. It's simple way for creating Kubernetes object files. (Deployment, Service, ConfigMap etc.) and It's the best way to find, share, and use software built for Kubernetes.

[Google Kubernetes Engine (GKE)](https://cloud.google.com/kubernetes-engine) provides a managed environment for deploying, managing, and scaling your containerized applications using Google infrastructure. The GKE environment consists of multiple machines (specifically, [Compute Engine](https://cloud.google.com/compute) instances) grouped together to form a [cluster](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture).

## Prerequisites

### Install Go tools

on **Linux**:
```bash
curl -L0 https://dl.google.com/go/go1.13.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.13.darwin-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

go version
go version go1.13 linux/amd64
```

on **macOS**:
```bash
curl -L0 https://dl.google.com/go/go1.13.darwin-amd64.tar.gz
tar -C /usr/local -xzf go1.13.darwin-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

go version
go version go1.13 darwin/amd64
```

on **Windows**:

1. [Download the zip file](https://dl.google.com/go/go1.13.windows-amd64.zip) and extract it into the directory of your choice (we suggest `c:\Go`).

2. Add the bin subdirectory of your Go root (for example, `c:\Go\bin`) to your `PATH` environment variable.

3. Setting environment variables under Windows Under Windows, you may set environment variables through the "Environment Variables" button on the "Advanced" tab of the "System" control panel. Some versions of Windows provide this control panel through the "Advanced System Settings" option inside the "System" control panel.

```bash
go version
go version go1.13 windows/amd64
```

### GitLab Variables

#### Service Account Key

The **Service Account** named **infra** must be created and **Service Account Key** must be added to **GitLab TodoListApplication Group Variables** as `INFRA_ADMIN_CREDENTIALS`.

#### Docker Credentials

The Docker Hub credentials must be added to **GitLab TodoListApplication Group Variables** as `DOCKER_USERNAME` and `DOCKER_PASSWORD` for `alicankustemur/go-echo` Docker image and install stage.

#### Other Variables
`PROJECT`, `REGION`, `BACKEND_EXTERNAL_IP` must be added to **GitLab TodoListApplication Group Variables**.

## Deployment

This repo has a pipeline. **(.gitlab-ci.yml)** 

The pipeline has four stages.

1. **install**
    - `build-go-echo-docker-image`
        - It creates a Go Echo Docker image with downloaded dependencies for speed up to `test` stage. It only runs if changes these files `Dockerfile-go.echo`, `go.mod` and `go.sum`.
        - This job uses `docker:19.03.8` image because only build a Docker image and it has one service named `docker:19.03.8-dind`. It waits the service is up&run then it builds docker image named `alicankustemur/go-echo` and push to Docker Hub.

2. **test**
    - `test`
        - It runs tests on `alicankustemur/go-echo` named Docker image.
        - It uses cache step, upload&download `/apt-cache` and `/go` paths for speed up.
        - Tests runs as parallel five test, shows coverage and result informations.

3. **build**
    - `build`
        - If `test` stage is success the job be triggered or fail the job be skipped.
        - It runs on `alicankustemur/gcloud-terraform` named Docker image because it needs gcloud [Google Container Registry (GCR)](https://cloud.google.com/container-registry) configuration.
        - This job uses a service named `docker:19.03.8-dind`. It waits the service is up&run then it builds docker image named for example `eu.gcr.io/todolist-app-1/back-end:2745e41a` and push to GCR.

4. **deploy**
    - `deploy`
        - If `build` stage is success the job be triggered or fail the job be skipped.
        - It runs on `alicankustemur/gcloud-terraform` named Docker image because it needs **GKE cluster** `KUBECONFIG` configurations.
        - `RELEASE_VERSION` variable is short commit id. Every commit is a version.
        - `DEPLOYED` variable does list helm packages and filter by `back-end`. If it returns 1 that means deployed a back-end container.
        - `sed` line replace **appVersion** value with `RELEASE_VERSION` on `Chart.yaml`.
        - If it deployed a `back-end` container it runs helm upgrade or it runs helm install.

#### Pipeline

It has only one docker image named `alicankustemur/gcloud-terraform`.
The docker image contains `gcloud-sdk`, `terraform`, `kubectl` and `helm` binaries.
The docker image is managed from [infrastructure/create-gcloud-terraform-docker-image](https://gitlab.com/todo-list-application/infrastructure/create-gcloud-terraform-docker-image) repository.

`.gcloud_login_before_script_template` only runs on `build` and `deploy` stage.
It creates a infra-admin-credentials.json file using `INFRA_ADMIN_CREDENTIALS` variable then authenticate to gcloud and the project with this infra-admin-credentials.json.

## Testing

This application has the following [handler tests.](https://echo.labstack.com/guide/testing)

[`handlers_test.go`](https://gitlab.com/todo-list-application/back-end/-/blob/master/src/todo/handlers_test.go)

```bash
TestGetAllTodos
TestCreateTodo
TestCreateTodoError
TestDeleteAllTodos
TestGetTodo
TestGetTodoError
TestGetTodoNotFound
TestDeleteTodo
TestDeleteTodoError
TestUpdateTodo
TestUpdateTodoError

coverage: 72.4% of statements
```



## Helm

Helm uses a packaging format called charts. A chart is a collection of files that describe a related set of Kubernetes resources. 

The `.helm` folder structure is:

```bash
.
├── Chart.yaml
├── charts
├── templates
│   ├── NOTES.txt
│   ├── _helpers.tpl
│   ├── deployment.yaml
│   └── service.yaml
└── values.yaml

2 directories, 6 files
```

`values.yaml` is values for `back-end` named **Helm Chart.**

> `values.yaml`

```yaml
# Default values for back-end.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: eu.gcr.io/todolist-app-1/back-end
  pullPolicy: IfNotPresent

imagePullSecrets: [ 
  {
    name: todolist-app-1-gke
  }
]

livenessProbePath: /todos
readinessProbePath: /todos

service:
  name: backend-end
  externalPort: 80
  internalPort: 8000
  type: LoadBalancer
  loadBalancer:
    loadBalancerIP: ""

  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
resources:
  limits:
    cpu: 100m
    memory: 128Mi
  requests:
    cpu: 100m
    memory: 128Mi

ingress:
  enabled: false
```

Every chart must have a version number. This is empty right now. But `deploy` stage on the pipeline creates a version using **Short Commit ID.**

> `Chart.yaml`

```yaml
apiVersion: v2
name: back-end
description: A Helm chart for Kubernetes

# A chart can be either an 'application' or a 'library' chart.
#
# Application charts are a collection of templates that can be packaged into versioned archives
# to be deployed.
#
# Library charts provide useful utilities or functions for the chart developer. They're included as
# a dependency of application charts to inject those utilities and functions into the rendering
# pipeline. Library charts do not define any templates and therefore cannot be deployed.
type: application

# This is the chart version. This version number should be incremented each time you make changes
# to the chart and its templates, including the app version.
version: 0.1.0

# This is the version number of the application being deployed. This version number should be
# incremented each time you make changes to the application.
appVersion: 
```

The Kubernetes Manifests, Deployment and Service are created by Helm using this files and values.

## Dockerfile

It has [multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/).

The `BUILDER` stage only fetchs dependencies and builds go application then creates a application binary under to `/build/todo`.

The `RUNTIME` stage only copies the created binary from `BUILDER` stage on `/build/todo` folder then runs the application named `todo`. Because of it is very small alpine image.

```Dockerfile
# Create build stage, this stage will just build and create "todo" named a binary file.
FROM golang:1.13-alpine as BUILDER

# Add git for fetch dependencies
RUN apk add \
        ca-certificates \
        git \
        build-base \
        --no-cache

WORKDIR /build/

# Copy Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY src/todo ./
RUN go build -v -o /build/todo

# Create runtime stage, this stage will just run builded "todo" binary file
FROM alpine:3.9 as RUNTIME

# This is important because application will not work.
# It gives this error : Env var 'PORT' must be set
ENV PORT 8000

WORKDIR /root

# Copy "todo" binary file from BUILDER stage
COPY --from=BUILDER /build/todo .

EXPOSE $PORT

ENTRYPOINT ["./todo"]
```

# Final

If you deployed this application to do cluster, Let's check this.

```
helm ls | grep back-end                                                                                                                                           
back-end	default  	1       	2020-06-10 22:43:00.301228657 +0000 UTC	deployed	back-end-0.1.0	83aa420c

kubectl get pods                                                                                 

NAME                        READY   STATUS    RESTARTS   AGE
back-end-5b9f847745-jmtkw   1/1     Running   0          3m5s

kubectl get svc                                                                                  

NAME         TYPE           CLUSTER-IP     EXTERNAL-IP    PORT(S)        AGE
back-end     LoadBalancer   10.27.244.50   34.65.248.48   80:30678/TCP   9m13s
kubernetes   ClusterIP      10.27.240.1    <none>         443/TCP        143m


curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"title":"buy some milk","order":1,"completed":false}' \
    http://34.65.248.48/todos | jq .                                                     
{
  "id": 1,
  "title": "buy some milk",
  "order": 1,
  "completed": false,
  "url": "http://34.65.248.48/todos/1"
}


curl http://34.65.248.48/todos | jq .                                                                                     
[
  {
    "id": 1,
    "title": "buy some milk",
    "order": 1,
    "completed": false,
    "url": "http://34.65.248.48/todos/1"
  }
]
```

🎉 Finally, **back-end** application is ready.