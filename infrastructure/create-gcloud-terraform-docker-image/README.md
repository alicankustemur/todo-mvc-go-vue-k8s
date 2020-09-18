# Creating Gloud SDK and Terraform Docker Image

This repo triggered automatically every commit.

It has only two files; `Dockerfile` and `.gitlab-ci.yml`.

Pipeline has one stage, It name is `build`.

It runs on `docker:19.03.8` image and creates `docker:19.03.8-dind` service.


It build, login Docker Hub account and push the `alicankustemur/gcloud-terraform:latest` image to Docker Hub.

> `.gitlab-ci.yml`

```yaml
image: docker:19.03.8

services:
  - docker:19.03.8-dind

variables:
  IMAGE: "alicankustemur/gcloud-terraform:latest"

build:
  stage: build
  script:
      - docker build -t $IMAGE .
      - docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
      - docker push $IMAGE
```


Dockerfile uses official `google/cloud-sdk:294.0.0-alpine` image at `FROM`.

`TERRAFORM_VERSION` and `HELM_VERSION` have `ARG` and `ENV`. Then the arguments have hard-coded versions. If you want change this versions, you can override;

```
docker build --build-arg TERRAFORM_VERSION=0.12.26 -t alicankustemur/gcloud-terraform:latest .
```


> `Dockerfile`

```Dockerfile
FROM google/cloud-sdk:294.0.0-alpine

ARG TERRAFORM_VERSION=0.12.16
ENV TERRAFORM_VERSION $TERRAFORM_VERSION

# Install terraform
RUN curl -O https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && mv terraform /usr/bin \
    && rm -rf terraform_${TERRAFORM_VERSION}_linux_amd64.zip

# Install kubectl
RUN gcloud components install kubectl

ARG HELM_VERSION=v3.1.2
ENV HELM_VERSION $HELM_VERSION

# Install helm
RUN curl -0 https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz | tar xz \
    && mv linux-amd64/helm /bin/helm \
    && rm -rf linux-amd64
```