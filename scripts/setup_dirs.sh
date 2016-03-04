#!/bin/bash -e

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow

# Directories
configs_dir=$main_dir/configs
staging_dir=$main_dir/staging
executables_dir=$main_dir/bin
unix_pipes_dir=$main_dir/pipes
es_dir=$main_dir/elasticsearch
scripts_dir=$main_dir/scripts
init_dir=$main_dir/init

# Pipes
raw_events_pipe=$unix_pipes_dir/raw-events-pipe
enriched_pipe=$unix_pipes_dir/enriched-events-pipe
bad_1_pipe=$unix_pipes_dir/bad-1-pipe

###########################
# Setup Directories/Files #
###########################

mkdir -p $configs_dir
mkdir -p $staging_dir
mkdir -p $executables_dir
mkdir -p $unix_pipes_dir
mkdir -p $es_dir
mkdir -p $scripts_dir
mkdir -p $init_dir

mkfifo $raw_events_pipe
mkfifo $enriched_pipe
mkfifo $bad_1_pipe
