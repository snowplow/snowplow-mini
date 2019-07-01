#!/bin/sh

# Create Kibana index patterns
curl -X POST \
  http://localhost:5601/api/saved_objects/index-pattern/good \
  -H 'Content-Type: application/json' \
  -H 'kbn-xsrf: true' \
  -d '{
  "attributes": {
    "title": "good",
    "timeFieldName": "collector_tstamp"
  }
}'

curl -X POST \
  http://localhost:5601/api/saved_objects/index-pattern/bad \
  -H 'Content-Type: application/json' \
  -H 'kbn-xsrf: true' \
  -d '{
  "attributes": {
    "title": "bad",
    "timeFieldName": "failure_tstamp"
  }
}'

# Set `good` as default index pattern
curl -X POST \
  http://localhost:5601/api/kibana/settings/defaultIndex \
  -H "Content-Type: application/json" \
  -H "kbn-xsrf: true" \
  -d '{
  "value": "good"
}'
