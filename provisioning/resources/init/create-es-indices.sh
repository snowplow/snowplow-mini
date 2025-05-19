#!/bin/sh

ES_HOST="localhost:9200"
MAPPINGS_DIR="/home/ubuntu/snowplow/elasticsearch/mapping"
MAX_RETRIES=12      # Total wait time = MAX_RETRIES * SLEEP_INTERVAL
SLEEP_INTERVAL=5    # in seconds

log() {
  echo "[$(date +'%Y-%m-%d %H:%M:%S')] $*"
}

create_index() {
  index_name=$1
  mapping_file="$MAPPINGS_DIR/${index_name}-mapping.json"
  retries=0

  if [ ! -f "$mapping_file" ]; then
    log "Mapping file not found: $mapping_file"
    exit 1
  fi

  log "Attempting to create index: $index_name"

  until curl --output /dev/null --silent --fail \
    -H 'Content-Type: application/json' \
    -X PUT "$ES_HOST/$index_name" \
    -d @"$mapping_file"; do

    retries=$((retries + 1))
    if [ "$retries" -ge "$MAX_RETRIES" ]; then
      log "Failed to create index '$index_name' after $MAX_RETRIES attempts."
      exit 1
    fi

    log "Retry $retries/$MAX_RETRIES: Waiting to create $index_name index..."
    sleep "$SLEEP_INTERVAL"
  done

  log "Successfully created index: $index_name"
}

create_index "good"
create_index "bad"
