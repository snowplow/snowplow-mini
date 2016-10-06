#!/bin/bash

# Copyright (c) 2016 Snowplow Analytics Ltd. All rights reserved.
#
# This program is licensed to you under the Apache License Version 2.0, and
# you may not use this file except in compliance with the Apache License
# Version 2.0.  You may obtain a copy of the Apache License Version 2.0 at
# http://www.apache.org/licenses/LICENSE-2.0.
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the Apache License Version 2.0 is distributed on an "AS
# IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
# implied.  See the Apache License Version 2.0 for the specific language
# governing permissions and limitations there under.

# Script to upload all schemas in a folder to the Iglu repo
# Takes three arguments: the target repo host, the Super API key, and the input schema directory
# Uses PUT rather than POST, so existing schemas are overwritten

echo "================================="
echo "  Starting Iglu Server Uploader"
echo "---------------------------------"

if [ "$#" -ne 3 ]
then
  echo "ERROR: 3 arguments required, $# provided"
  exit 1
fi

host=$1;
apikey=$2;
schemafolder=$3;

echo ""
echo "Making all_vendor API Keys:"
api_keys="$(curl --silent ${host}/api/auth/keygen -X POST -H "apikey: ${apikey}" -d "vendor_prefix=*")"
write_api_key="$(echo ${api_keys} | python -c 'import json,sys;obj=json.load(sys.stdin);print(obj["write"])')"
read_api_key="$(echo ${api_keys} | python -c 'import json,sys;obj=json.load(sys.stdin);print(obj["read"])')"
echo "Keys: $(echo ${api_keys} | xargs)"

echo ""
echo "Uploading all Schemas found in ${schemafolder}:"

good_counter=0
bad_counter=0

for schemapath in $(find $schemafolder -type f | grep 'jsonschema'); do
  destination="$host/api/schemas/$(
    # Keep the last 4 slash-separated components of the filename
    echo $schemapath | awk -F '/' '{print $(NF-3)"/"$(NF-2)"/"$(NF-1)"/"$(NF)}';
  )";
  echo "Uploading schema in file '$schemapath' to endpoint '$destination'";
  result="$(curl --silent "${destination}?isPublic=true" -XPUT -d @$schemapath -H "apikey: $write_api_key")";
  echo " - Result: $(echo ${result} | xargs)"

  # Process result
  status="$(echo ${result} | python -c 'import json,sys;obj=json.load(sys.stdin);print(obj["status"])')"
  if [[ "${status}" -eq "200" ]] || [[ "${status}" -eq "201" ]]; then
    let good_counter=good_counter+1
  else
    let bad_counter=bad_counter+1
  fi
done;

echo ""
echo "Result Counts:"
echo " - 200/201: ${good_counter}"
echo " - 400/401/500: ${bad_counter}"

echo ""
echo "Remove created API Keys:"
echo " - Remove ${write_api_key}: $(curl --silent ${host}/api/auth/keygen -X DELETE -H "apikey: ${apikey}" -d "key=${write_api_key}" | xargs)"
echo " - Remove ${read_api_key}: $(curl --silent ${host}/api/auth/keygen -X DELETE -H "apikey: ${apikey}" -d "key=${read_api_key}" | xargs)"

echo ""
echo "--------"
echo "  Done  "
echo "========"
