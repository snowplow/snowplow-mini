#!/bin/bash -e

#sudo apt-get update
#sudo apt-get install -y unzip

#############
# Constants #
#############

main_dir=/home/ubuntu/snowplow

# Directories
staging_dir=$main_dir/staging
executables_dir=$main_dir/bin
es_dir=$main_dir/elasticsearch

# Packagaes
kinesis_package=snowplow_kinesis_r85_metamorphosis.zip
iglu_server_package=iglu_server_0.2.0.zip
kibana_v=4.0.1

##################
# Install Java 7 #
##################

#sudo add-apt-repository ppa:webupd8team/java -y
#sudo apt-get update
#echo oracle-java7-installer shared/accepted-oracle-license-v1-1 select true | sudo /usr/bin/debconf-set-selections
#sudo apt-get install oracle-java7-installer -y

##################
# Install GO 1.7 #
##################

wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz -P $staging_dir
sudo tar -C /usr/local -xzf ${staging_dir}/go1.7.3.linux-amd64.tar.gz
sudo mkdir -p ~/go
echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc


####################################################
# Build control-plane-api and install dependencies #
####################################################

export GOPATH=/home/vagrant/go
export PATH=$PATH":/usr/local/go/bin"

dir="/vagrant/control-plane"
cd $dir

go get github.com/emicklei/go-restful
go build -o snowplow-mini-control-plane-api snowplow-mini-control-plane-api.go
sudo cp snowplow-mini-control-plane-api $executables_dir


################################
# Install Kinesis Applications #
################################

#wget http://dl.bintray.com/snowplow/snowplow-generic/${kinesis_package} -P $staging_dir
#unzip $staging_dir/${kinesis_package} -d $executables_dir

#######################
# Install Iglu Server #
#######################

#wget http://bintray.com/artifact/download/snowplow/snowplow-generic/${iglu_server_package} -P $staging_dir
#unzip $staging_dir/${iglu_server_package} -d $executables_dir
#sudo -u postgres psql -c "create user snowplow createdb password 'snowplow';" || true
#sudo -u postgres psql -c "create database iglu owner snowplow;" || true

#########################
# Install Elasticsearch #
#########################

#wget -qO - https://packages.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
#echo "deb http://packages.elastic.co/elasticsearch/1.4/debian stable main" | sudo tee -a /etc/apt/sources.list
#sudo apt-get update -y && sudo apt-get install elasticsearch -y
#sudo /usr/share/elasticsearch/bin/plugin --install mobz/elasticsearch-head

##################
# Install Kibana #
##################

#wget "https://download.elasticsearch.org/kibana/kibana/kibana-${kibana_v}-linux-x64.zip" -P $staging_dir
#sudo unzip $staging_dir/kibana-${kibana_v}-linux-x64.zip -d /opt/
#sudo ln -s /opt/kibana-${kibana_v}-linux-x64 /opt/kibana

#################
# Install Nginx #
#################

#sudo apt-get -y install nginx apache2-dev

# Set ownership of directory
#sudo chown -R ubuntu:ubuntu $main_dir
