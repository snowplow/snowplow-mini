#!/bin/bash

# NOTE: Use `uuidgen` to create new `uid`
iglu_server_super_uid="980ae3ab-3aba-4ffe-a3c2-3b2e24e2ffce"

domain_name=example.com
tls_cond="off"

username=USERNAME_PLACEHOLDER
password=PASSWORD_PLACEHOLDER


# DO NOT ALTER BELOW #
#add apiKey to iglu-resolver.json for auth in the iglu server
iglu_resolver_config_dir="/home/ubuntu/snowplow/configs/iglu-resolver.json"
sed -i 's/\(.*"apikey":\)\(.*\)/\1 "'$iglu_server_super_uid'"/' $iglu_resolver_config_dir

#write super apikey to db
export PGPASSWORD=snowplow
iglu_server_setup="INSERT INTO apikeys (uid, vendor_prefix, permission, createdat) VALUES ('${iglu_server_super_uid}','*','super',current_timestamp);"
psql --host=localhost --port=5432 --username=snowplow --dbname=iglu -c "${iglu_server_setup}"

#add domain name to Caddyfile
inserted_line=""
sed -i '1d' /home/ubuntu/snowplow/configs/Caddyfile #delete first line of the default Caddyfile 
if [[ "${tls_cond}" == "on" ]]; then
  inserted_line="$domain_name *:80 { \n        tls example@example.com \n"
else
  inserted_line="*:80 { \n        tls off \n"
fi
sed -i "1s/^/${inserted_line}/" /home/ubuntu/snowplow/configs/Caddyfile

#add username and password to Caddyfile for basic auth
sed -i "s/USERNAME_PLACEHOLDER/$username/g" /home/ubuntu/snowplow/configs/Caddyfile
sed -i "s/PASSWORD_PLACEHOLDER/$password/g" /home/ubuntu/snowplow/configs/Caddyfile
sudo service caddy_init restart
