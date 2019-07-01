#!/bin/sh

# Create Elasticsearch indices
curl -H 'Content-Type: application/json' -X PUT localhost:9200/good -d @/home/ubuntu/snowplow/elasticsearch/mapping/good-mapping.json
curl -H 'Content-Type: application/json' -X PUT localhost:9200/bad -d @/home/ubuntu/snowplow/elasticsearch/mapping/bad-mapping.json
