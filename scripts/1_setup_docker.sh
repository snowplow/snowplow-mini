#!/bin/bash -e

apt-get update
apt-get install -y unzip

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow

configs_dir=$main_dir/configs
staging_dir=$main_dir/staging
executables_dir=$main_dir/bin
unix_pipes_dir=$main_dir/pipes
es_dir=$main_dir/elasticsearch
scripts_dir=$main_dir/scripts

raw_events_pipe=$unix_pipes_dir/raw-events-pipe
enriched_pipe=$unix_pipes_dir/enriched-events-pipe

kinesis_package=snowplow_kinesis_r67_bohemian_waxwing.zip
kibana_v=4.0.1

###########################
# Setup Directories/Files #
###########################

mkdir -p $configs_dir
mkdir -p $staging_dir
mkdir -p $executables_dir
mkdir -p $unix_pipes_dir
mkdir -p $es_dir
mkdir -p $scripts_dir

mkfifo $raw_events_pipe
mkfifo $enriched_pipe

################################
# Install Kinesis Applications #
################################

wget http://dl.bintray.com/snowplow/snowplow-generic/${kinesis_package} -P $staging_dir
unzip $staging_dir/${kinesis_package} -d $executables_dir

#########################
# Install Elasticsearch #
#########################

wget -qO - https://packages.elastic.co/GPG-KEY-elasticsearch | apt-key add -
echo "deb http://packages.elastic.co/elasticsearch/1.4/debian stable main" | tee -a /etc/apt/sources.list
apt-get update -y && apt-get install elasticsearch -y
/usr/share/elasticsearch/bin/plugin --install mobz/elasticsearch-head

##################
# Install Kibana #
##################

wget "https://download.elasticsearch.org/kibana/kibana/kibana-${kibana_v}-linux-x64.zip" -P $staging_dir
unzip $staging_dir/kibana-${kibana_v}-linux-x64.zip -d /opt/
ln -s /opt/kibana-${kibana_v}-linux-x64 /opt/kibana
