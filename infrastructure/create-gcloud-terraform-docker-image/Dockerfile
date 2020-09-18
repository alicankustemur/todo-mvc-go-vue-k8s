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