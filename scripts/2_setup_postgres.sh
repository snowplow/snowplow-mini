#!/bin/bash -e

# NOTE: Seperated from provision.sh as Travis comes with Postgres

sudo apt-get update
sudo apt-get install -y unzip

######################
# Install PostgreSQL #
######################

sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" >> /etc/apt/sources.list.d/pgdg.list'
wget -q https://www.postgresql.org/media/keys/ACCC4CF8.asc -O - | sudo apt-key add -
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib -y
