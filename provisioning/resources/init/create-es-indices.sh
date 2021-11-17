#!/bin/sh

# Create Elasticsearch indices
until $(curl --output /dev/null --silent --fail -H 'Content-Type: application/json' -X PUT localhost:9200/good -d @/home/ubuntu/snowplow/elasticsearch/mapping/good-mapping.json); do echo "creating good index"; sleep 5; done
until $(curl --output /dev/null --silent --fail -H 'Content-Type: application/json' -X PUT localhost:9200/bad -d @/home/ubuntu/snowplow/elasticsearch/mapping/bad-mapping.json); do echo "creating bad index"; sleep 5; done
