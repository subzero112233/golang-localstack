#!/usr/bin/env bash

set -euo pipefail

echo "configuring aws infrastructure"
echo "==================="
export LOCALSTACK_HOST=localhost
export AWS_DEFAULT_REGION=eu-central-1
export AWS_ACCESS_KEY_ID=irrelevant
export AWS_SECRET_ACCESS_KEY=irrelevant


create_queue() {
    local QUEUE_NAME_TO_CREATE=$1
    aws --endpoint-url=http://${LOCALSTACK_HOST}:4566 sqs create-queue --queue-name ${QUEUE_NAME_TO_CREATE} --attributes VisibilityTimeout=30
}

create_bucket() {
    local BUCKET_NAME_TO_CREATE=$1
    aws --endpoint-url=http://${LOCALSTACK_HOST}:4566 s3 mb s3://${BUCKET_NAME_TO_CREATE}
}

# create queues
for queue in "queue1"
do
  create_queue $queue
done


# create buckets
for bucket in "bucket1"
do
  create_bucket $bucket
done

echo "finished running localstack bootstrap"
