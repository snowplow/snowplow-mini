#!/bin/bash
set -e

tag=$1

tag_version=`echo $tag | cut -d'/' -f1`
tag_platform=`echo $tag | cut -d'/' -f2`
tag_flavor=`echo $tag | cut -d'/' -f3`

version=`cat ${TRAVIS_BUILD_DIR}/VERSION`

if [ "${tag_version}" == "${version}" ]; then
  if [ "${tag_platform}" == "aws" ] || [ "${tag_platform}" == "gcp" ]; then
    if [ "${tag_flavor}" == "large" ] || [ "${tag_flavor}" == "xlarge" ] || [ "${tag_flavor}" == "xxlarge" ]; then
      exit 0
    fi
    echo "Tag flavor is not valid!"
    exit 1
  fi
  echo "Tag platform is not valid!"
  exit 1
else
  echo "Tag version doesn't match VERSION file!"
  exit 1
fi
