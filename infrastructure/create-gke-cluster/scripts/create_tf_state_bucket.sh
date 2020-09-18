# !/bin/bash

PROJECT=$1
SERVICE_ACCOUNT=$2
REGION=$3
BUCKET_NAME=$4

gsutil mb \
    -p $PROJECT \
    -c regional \
    -l $REGION \
    gs://$BUCKET_NAME/

gsutil versioning set on gs://$BUCKET_NAME/

gsutil iam ch serviceAccount:$SERVICE_ACCOUNT@$PROJECT.iam.gserviceaccount.com:legacyBucketWriter gs://$BUCKET_NAME/