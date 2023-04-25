#!/bin/sh

# Create Kibana index patterns

until $(curl --output /dev/null --silent --fail -H 'Content-Type: application/json' -H 'osd-xsrf: true' -X POST http://localhost:5601/api/saved_objects/index-pattern/good -d '{"attributes": {"title": "good","timeFieldName": "collector_tstamp"}}'); do echo "creating good index pattern"; sleep 5; done

until $(curl --output /dev/null --silent --fail -H 'Content-Type: application/json' -H 'osd-xsrf: true' -X POST http://localhost:5601/api/saved_objects/index-pattern/bad -d '{"attributes": {"title": "bad","timeFieldName": "data.failure.timestamp"}}'); do echo "creating bad index pattern"; sleep 5; done

# Set `good` as default index pattern
until $(curl --output /dev/null --silent --fail -H 'Content-Type: application/json' -H 'osd-xsrf: true' -X POST http://localhost:5601/api/opensearch-dashboards/settings/defaultIndex -d '{"value": "good"}'); do echo "setting [good] as the default index pattern"; sleep 5; done
