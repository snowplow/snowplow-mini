#!/bin/bash -e

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow
configs_dir=$main_dir/configs

############
# Setup UI #
############

sudo rm -f /etc/nginx/sites-enabled/default
sudo cp $configs_dir/snowplow-mini.conf /etc/nginx/conf.d/

