#!/bin/bash -e

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow

configs_dir=$main_dir/configs
executables_dir=$main_dir/bin
unix_pipes_dir=$main_dir/pipes
es_dir=$main_dir/elasticsearch
scripts_dir=$main_dir/scripts

raw_events_pipe=$unix_pipes_dir/raw-events-pipe
enriched_pipe=$unix_pipes_dir/enriched-events-pipe

#####################
# Start ES + Kibana #
#####################

service elasticsearch start
service kibana4_init start

sleep 15

################
# Add Mappings #
################

curl -XPUT 'http://localhost:9200/good' -d @${es_dir}/good-mapping.json
curl -XPUT 'http://localhost:9200/bad' -d @${es_dir}/bad-mapping.json

#################################################
# Start Collector/Enrichment/Elasticsearch Sink #
#################################################

${executables_dir}/snowplow-stream-collector-0.5.0 --config ${configs_dir}/scala-stream-collector.hocon > $raw_events_pipe &
cat $raw_events_pipe | ${executables_dir}/snowplow-kinesis-enrich-0.6.0 --config ${configs_dir}/scala-kinesis-enrich.hocon --resolver file:${configs_dir}/default-iglu-resolver.json > $enriched_pipe &
cat $enriched_pipe | ${executables_dir}/snowplow-elasticsearch-sink-0.4.0 --config ${configs_dir}/kinesis-elasticsearch-sink-good.hocon | ${scripts_dir}/elasticsearch_upload.pl
