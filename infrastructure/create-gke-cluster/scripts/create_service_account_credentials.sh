# !/bin/bash

PROJECT=$1
SERVICE_ACCOUNT=$2
SERVICE_ACCOUNT_CREDENTIALS=$3

ROLES=(
container.admin
compute.admin
iam.serviceAccountAdmin
iam.serviceAccountUser
resourcemanager.projectIamAdmin
serviceusage.serviceUsageAdmin
storage.objectAdmin
iam.serviceAccountKeyAdmin
compute.publicIpAdmin
compute.securityAdmin
)

gcloud iam service-accounts create $SERVICE_ACCOUNT --display-name $SERVICE_ACCOUNT

for ROLE in "${ROLES[@]}"; do
        gcloud projects add-iam-policy-binding $PROJECT --member serviceAccount:$SERVICE_ACCOUNT@$PROJECT.iam.gserviceaccount.com --role roles/$ROLE
done

gcloud iam service-accounts keys create $SERVICE_ACCOUNT_CREDENTIALS \
        --iam-account=$SERVICE_ACCOUNT@$PROJECT.iam.gserviceaccount.com