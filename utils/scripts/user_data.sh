#!/bin/bash

set -e -x

domain_name=example.com
username='username'
password='password'
iglu_server_super_uid='deadbeef-dead-beef-dead-beefdeadbeef'

# DO NOT ALTER BELOW #
docker compose -f /home/ubuntu/snowplow/docker-compose.yml restart iglu-server
sudo service snowplow_mini_control_plane_api restart

sleep 10

# Add domain name to Caddyfile
curl -XPOST -d "domain_name=$domain_name" localhost:10000/domain-name

# Add username and password to Caddyfile for basic auth
curl -XPOST -d "new_username=$username&new_password=$password" localhost:10000/credentials

# Add apiKey to iglu-resolver.json for auth in the iglu server
curl -XPOST -d "local_iglu_apikey=$iglu_server_super_uid" localhost:10000/local-iglu-apikey
