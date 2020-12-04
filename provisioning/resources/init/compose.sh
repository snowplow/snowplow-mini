#!/bin/sh

platform=$(cat /home/ubuntu/snowplow/configs/compositions/.platform)

if [ "$platform" = "aws" ]; then
  cat /home/ubuntu/snowplow/configs/compositions/docker-compose-aws.yml > /home/ubuntu/snowplow/docker-compose.yml
fi

if [ "$platform" = "gcp" ]; then
  cat /home/ubuntu/snowplow/configs/compositions/docker-compose-gcp.yml > /home/ubuntu/snowplow/docker-compose.yml
fi
