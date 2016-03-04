#!/bin/bash

# NOTE: Use `uuidgen` to create new `uid`
iglu_server_super_uid="980ae3ab-3aba-4ffe-a3c2-3b2e24e2ffce"

# DO NOT ALTER BELOW #

export PGPASSWORD=snowplow
iglu_server_setup="INSERT INTO apikeys (uid, vendor_prefix, permission, createdat) VALUES ('${iglu_server_super_uid}','*','super',current_timestamp);"
psql --host=localhost --port=5432 --username=snowplow --dbname=iglu -c "${iglu_server_setup}"
