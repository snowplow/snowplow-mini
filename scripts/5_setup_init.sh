#!/bin/bash -e

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow
init_dir=$main_dir/init

##########################
# Setup all init.d files #
##########################

sudo update-rc.d elasticsearch defaults

sudo cp $init_dir/kibana4_init /etc/init.d
sudo chmod 0755 /etc/init.d/kibana4_init
sudo update-rc.d kibana4_init defaults

sudo cp $init_dir/snowplow_stream_collector_0.9.0 /etc/init.d
sudo chmod 0755 /etc/init.d/snowplow_stream_collector_0.9.0
sudo update-rc.d snowplow_stream_collector_0.9.0 defaults

sudo cp $init_dir/snowplow_stream_enrich_0.10.0 /etc/init.d
sudo chmod 0755 /etc/init.d/snowplow_stream_enrich_0.10.0
sudo update-rc.d snowplow_stream_enrich_0.10.0 defaults

sudo cp $init_dir/snowplow_elasticsearch_sink_good_0.8.0 /etc/init.d
sudo chmod 0755 /etc/init.d/snowplow_elasticsearch_sink_good_0.8.0
sudo update-rc.d snowplow_elasticsearch_sink_good_0.8.0 defaults

sudo cp $init_dir/snowplow_elasticsearch_sink_bad_0.8.0 /etc/init.d
sudo chmod 0755 /etc/init.d/snowplow_elasticsearch_sink_bad_0.8.0
sudo update-rc.d snowplow_elasticsearch_sink_bad_0.8.0 defaults

sudo cp $init_dir/iglu_server_0.2.0 /etc/init.d
sudo chmod 0755 /etc/init.d/iglu_server_0.2.0
sudo update-rc.d iglu_server_0.2.0 defaults

sudo cp $init_dir/nginx_passenger /etc/init.d
sudo chmod 0755 /etc/init.d/nginx_passenger
sudo update-rc.d nginx_passenger defaults
