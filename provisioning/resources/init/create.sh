#!/bin/sh

# Create Elasticsearch indices
curl -H 'Content-Type: application/json' -X PUT localhost:9200/good -d @/home/ubuntu/snowplow/elasticsearch/mapping/good-mapping.json && \
curl -H 'Content-Type: application/json' -X PUT localhost:9200/bad -d @/home/ubuntu/snowplow/elasticsearch/mapping/bad-mapping.json && \

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

# Create NSQ topics
curl -X POST localhost:4151/topic/create?topic=RawEvents && \
curl -X POST localhost:4151/topic/create?topic=BadEvents && \
curl -X POST localhost:4151/topic/create?topic=EnrichedEvents && \
curl -X POST localhost:4151/topic/create?topic=BadEnrichedEvents
