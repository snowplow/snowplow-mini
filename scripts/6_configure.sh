#!/bin/bash -e

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow
es_dir=$main_dir/elasticsearch

################
# Add Mappings #
################

sudo service elasticsearch start
sleep 15

curl -XPUT 'http://localhost:9200/good' -d @${es_dir}/good-mapping.json
curl -XPUT 'http://localhost:9200/bad' -d @${es_dir}/bad-mapping.json

####################
# Init Iglu Server #
####################

sudo service iglu_server_0.2.0 start
sleep 30
