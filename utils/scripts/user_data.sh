#!/bin/bash

# NOTE: Use `uuidgen` to create new `uid`
iglu_server_super_uid="980ae3ab-3aba-4ffe-a3c2-3b2e24e2ffce"

domain_name=example.com
tls_cond="off"

# DO NOT ALTER BELOW #
iglu_resolver_config_dir="/home/ubuntu/snowplow/configs/iglu-resolver.json"
sed -i 's/\(.*"apikey":\)\(.*\)/\1 "'$iglu_server_super_uid'"/' $iglu_resolver_config_dir

export PGPASSWORD=snowplow
iglu_server_setup="INSERT INTO apikeys (uid, vendor_prefix, permission, createdat) VALUES ('${iglu_server_super_uid}','*','super',current_timestamp);"
psql --host=localhost --port=5432 --username=snowplow --dbname=iglu -c "${iglu_server_setup}"

inserted_line=""

sed -i '1d' /home/ubuntu/snowplow/configs/Caddyfile #delete first line of the default Caddyfile 

if [[ "${tls_cond}" == "on" ]]; then
  inserted_line="$domain_name *:80 { \n        tls example@example.com \n"
else
  inserted_line="*:80 { \n        tls off \n"
fi


sed -i "1s/^/${inserted_line}/" /home/ubuntu/snowplow/configs/Caddyfile
