#!/bin/bash -e

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow
staging_dir=$main_dir/staging

############################
# Remove Staging Directory #
############################

rm -rf $staging_dir
