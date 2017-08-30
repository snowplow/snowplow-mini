#!/bin/bash

sudo service elasticsearch start
sudo service iglu_server_0.2.0 start
sudo service snowplow_stream_collector_0.9.0 start
sudo service snowplow_stream_enrich_0.10.0 start
sudo service snowplow_elasticsearch_sink_good_0.8.0 start
sudo service snowplow_elasticsearch_sink_bad_0.8.0 start
sudo service kibana4_init start
sudo service nginx start
sleep 15

# Send good and bad events
COUNTER=0
while [  $COUNTER -lt 10 ]; do
  curl http://localhost:8080/i?e=pv
  curl http://localhost:8080/i
  let COUNTER=COUNTER+1 
done
sleep 5

# Assertions
good_count="$(curl --silent -XGET 'http://localhost:9200/good/good/_count' | python -c 'import json,sys;obj=json.load(sys.stdin);print obj["count"]')"
bad_count="$(curl --silent -XGET 'http://localhost:9200/bad/bad/_count' | python -c 'import json,sys;obj=json.load(sys.stdin);print obj["count"]')"

echo "Event Counts:"
echo " - Good: ${good_count}"
echo " - Bad: ${bad_count}"

stream_enrich_pid_file=/var/run/snowplow_stream_enrich_0.10.0.pid
stream_collector_pid_file=/var/run/snowplow_stream_collector_0.9.0.pid
sink_bad_pid_file=/var/run/snowplow_elasticsearch_sink_bad_0.8.0-2x.pid
sink_good_pid_file=/var/run/snowplow_elasticsearch_sink_good_0.8.0-2x.pid


stream_enrich_pid_old="$(cat "${stream_enrich_pid_file}")"
stream_collector_pid_old="$(cat "${stream_collector_pid_file}")"
sink_bad_pid_old="$(cat "${sink_bad_pid_file}")"
sink_good_pid_old="$(cat "${sink_good_pid_file}")"

req_result=$(curl --silent -XPUT 'http://localhost:10000/restart-services')

stream_enrich_pid_new="$(cat "${stream_enrich_pid_file}")"
stream_collector_pid_new="$(cat "${stream_collector_pid_file}")"
sink_bad_pid_new="$(cat "${sink_bad_pid_file}")"
sink_good_pid_new="$(cat "${sink_good_pid_file}")"

# Bad Count is 11 due to bad logging
if [[ "${good_count}" -eq "10" ]] && [[ "${bad_count}" -eq "11" ]] && 
   [[ "${req_result}" == "OK" ]] &&
   [[ "${stream_enrich_pid_old}" -ne "${stream_enrich_pid_new}" ]] &&
   [[ "${stream_collector_pid_old}" -ne "${stream_collector_pid_new}" ]] &&
   [[ "${sink_bad_pid_old}" -ne "${sink_bad_pid_new}" ]] &&
   [[ "${sink_good_pid_old}" -ne "${sink_good_pid_new}" ]]; then

  exit 0
else
  exit 1
fi
