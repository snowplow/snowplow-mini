#!/bin/bash

# Send good and bad events, and test that the default redirect path is disabled
COUNTER=0
while [  $COUNTER -lt 10 ]; do
  curl http://localhost:8080/i?e=pv
  curl http://localhost:8080/i
  curl http://localhost:8080/r/tp2
  let COUNTER=COUNTER+1
done
sleep 30

# Assertions
good_count="$(curl --silent -XGET 'http://localhost:9200/good/good/_count' | python -c 'import json,sys;obj=json.load(sys.stdin);print obj["count"]')"
bad_count="$(curl --silent -XGET 'http://localhost:9200/bad/bad/_count' | python -c 'import json,sys;obj=json.load(sys.stdin);print obj["count"]')"

echo "Event Counts:"
echo " - Good: ${good_count}"
echo " - Bad: ${bad_count}"

if [[ "${good_count}" -eq "10" ]] && [[ "${bad_count}" -eq "10" ]]; then
  exit 0
else
  exit 1
fi
