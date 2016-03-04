#!/bin/bash -e

sudo apt-get update
sudo apt-get install -y unzip

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow

# Directories
staging_dir=$main_dir/staging
executables_dir=$main_dir/bin
es_dir=$main_dir/elasticsearch

# Packagaes
kinesis_package=snowplow_kinesis_r78_great_hornbill.zip
kibana_v=4.0.1

##################
# Install Java 7 #
##################

sudo add-apt-repository ppa:webupd8team/java -y
sudo apt-get update
echo oracle-java7-installer shared/accepted-oracle-license-v1-1 select true | sudo /usr/bin/debconf-set-selections
sudo apt-get install oracle-java7-installer -y

################################
# Install Kinesis Applications #
################################

wget http://dl.bintray.com/snowplow/snowplow-generic/${kinesis_package} -P $staging_dir
unzip $staging_dir/${kinesis_package} -d $executables_dir

#########################
# Install Elasticsearch #
#########################

wget -qO - https://packages.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
echo "deb http://packages.elastic.co/elasticsearch/1.4/debian stable main" | sudo tee -a /etc/apt/sources.list
sudo apt-get update -y && sudo apt-get install elasticsearch -y
sudo /usr/share/elasticsearch/bin/plugin --install mobz/elasticsearch-head

##################
# Install Kibana #
##################

wget "https://download.elasticsearch.org/kibana/kibana/kibana-${kibana_v}-linux-x64.zip" -P $staging_dir
sudo unzip $staging_dir/kibana-${kibana_v}-linux-x64.zip -d /opt/
sudo ln -s /opt/kibana-${kibana_v}-linux-x64 /opt/kibana

sudo chown -R ubuntu:ubuntu $main_dir

################
# Add Mappings #
################

sudo service elasticsearch start
sleep 15

curl -XPUT 'http://localhost:9200/good' -d @${es_dir}/good-mapping.json
curl -XPUT 'http://localhost:9200/bad' -d @${es_dir}/bad-mapping.json
