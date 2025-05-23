#!/bin/bash
set -e

cd $GITHUB_WORKSPACE

###### input validation ######
if [ "$#" -ne 1 ]
then
	echo "Inappropriate number of arguments! Aborted."
	echo "Usage: deploy.sh tag"
	exit 1
else
	export SP_MINI_VERSION=`echo $1 | cut -d'/' -f1`
	export SP_MINI_PLATFORM=`echo $1 | cut -d'/' -f2`
	export SP_MINI_SIZE=`echo $1 | cut -d'/' -f3`

	if [ "$SP_MINI_PLATFORM" == "aws" -o "$SP_MINI_PLATFORM" == "gcp" ]
	then
		echo "Platform recognized!"
	else
		echo "Unrecognized platform! Aborted."
		echo "Supported platforms: aws, gcp."
		exit 1
	fi

	if [ "$SP_MINI_SIZE" == "large" ]
	then
		# prepare env vars for docker compose
		echo -n > provisioning/roles/docker/files/.env
		echo "ELASTICSEARCH_MEM_SIZE=2560m" >> provisioning/roles/docker/files/.env
		echo "ELASTICSEARCH_HEAP_SIZE=1280m" >> provisioning/roles/docker/files/.env
		echo "KIBANA_MEM_SIZE=512m" >> provisioning/roles/docker/files/.env
		echo "SP_COLLECTOR_MEM_SIZE=512m" >> provisioning/roles/docker/files/.env
		echo "SP_ENRICH_MEM_SIZE=1g" >> provisioning/roles/docker/files/.env
		echo "SP_ES_LOADER_MEM_SIZE=512m" >> provisioning/roles/docker/files/.env
		echo "SP_IGLU_SERVER_MEM_SIZE=386m" >> provisioning/roles/docker/files/.env
		echo "CA_ADVISOR_MEM_SIZE=128m" >> provisioning/roles/docker/files/.env
		echo "NSQD_MEM_SIZE=256m" >> provisioning/roles/docker/files/.env
		# prepare env var for packer
		# to be used to determine which instance type to use
		if [ "$SP_MINI_PLATFORM" == "aws" ]
		then
			export AWS_INSTANCE_TYPE=t2.large
		elif [ "$SP_MINI_PLATFORM" == "gcp" ]
		then
			export GCP_MACHINE_TYPE=n1-standard-2
		fi
	elif [ "$SP_MINI_SIZE" == "xlarge" ]
	then
		# prepare env vars for docker compose
		echo -n > provisioning/roles/docker/files/.env
		echo "ELASTICSEARCH_MEM_SIZE=5g" >> provisioning/roles/docker/files/.env
		echo "ELASTICSEARCH_HEAP_SIZE=2560m" >> provisioning/roles/docker/files/.env
		echo "KIBANA_MEM_SIZE=1g" >> provisioning/roles/docker/files/.env
		echo "SP_COLLECTOR_MEM_SIZE=1g" >> provisioning/roles/docker/files/.env
		echo "SP_ENRICH_MEM_SIZE=2g" >> provisioning/roles/docker/files/.env
		echo "SP_ES_LOADER_MEM_SIZE=1g" >> provisioning/roles/docker/files/.env
		echo "SP_IGLU_SERVER_MEM_SIZE=512m" >> provisioning/roles/docker/files/.env
		echo "CA_ADVISOR_MEM_SIZE=256m" >> provisioning/roles/docker/files/.env
		echo "NSQD_MEM_SIZE=512m" >> provisioning/roles/docker/files/.env
		# prepare env var for packer
		# to be used to determine which instance type to use
		if [ "$SP_MINI_PLATFORM" == "aws" ]
		then
			export AWS_INSTANCE_TYPE=t2.xlarge
		elif [ "$SP_MINI_PLATFORM" == "gcp" ]
		then
			export GCP_MACHINE_TYPE=n1-standard-4
		fi
	elif [ "$SP_MINI_SIZE" == "xxlarge" ]
	then
		# prepare env vars for docker compose
		echo -n > provisioning/roles/docker/files/.env
		echo "ELASTICSEARCH_MEM_SIZE=10g" >> provisioning/roles/docker/files/.env
		echo "ELASTICSEARCH_HEAP_SIZE=5g" >> provisioning/roles/docker/files/.env
		echo "KIBANA_MEM_SIZE=2g" >> provisioning/roles/docker/files/.env
		echo "SP_COLLECTOR_MEM_SIZE=2g" >> provisioning/roles/docker/files/.env
		echo "SP_ENRICH_MEM_SIZE=4g" >> provisioning/roles/docker/files/.env
		echo "SP_ES_LOADER_MEM_SIZE=2g" >> provisioning/roles/docker/files/.env
		echo "SP_IGLU_SERVER_MEM_SIZE=512m" >> provisioning/roles/docker/files/.env
		echo "CA_ADVISOR_MEM_SIZE=512m" >> provisioning/roles/docker/files/.env
		echo "NSQD_MEM_SIZE=1g" >> provisioning/roles/docker/files/.env
		# prepare env var for packer
		# to be used to determine which instance type to use
		if [ "$SP_MINI_PLATFORM" == "aws" ]
		then
			export AWS_INSTANCE_TYPE=t2.2xlarge
		elif [ "$SP_MINI_PLATFORM" == "gcp" ]
		then
			export GCP_MACHINE_TYPE=n1-standard-8
		fi
	else
		echo "Unrecognized size! Aborted."
		echo "Available sizes: large, xlarge and xxlarge."
		exit 1
	fi
fi

############## deployment ###################

if [ "$SP_MINI_PLATFORM" == "aws" ]
then
	echo -n > provisioning/resources/configs/compositions/.platform
	echo "aws" >> provisioning/resources/configs/compositions/.platform
	echo "AWS_LOGS_GROUP=snowplow-mini" >> provisioning/roles/docker/files/.env
	export AWS_ACCESS_KEY_ID=$AWS_DEPLOY_ACCESS_KEY
	export AWS_SECRET_ACCESS_KEY=$AWS_DEPLOY_SECRET_KEY
	export AWS_SP_MINI_VERSION=$SP_MINI_VERSION
	packer build -only=amazon-ebs Packerfile.json
elif [ "$SP_MINI_PLATFORM" == "gcp" ]
then
	echo -n > provisioning/resources/configs/compositions/.platform
	echo "gcp" >> provisioning/resources/configs/compositions/.platform
	echo $GOOGLE_APPLICATION_CREDENTIALS_JSON_BASE64 | base64 --decode > $GITHUB_WORKSPACE/BIN
	export GOOGLE_APPLICATION_CREDENTIALS=$GITHUB_WORKSPACE/BIN
	export GCP_SP_MINI_VERSION=`echo $SP_MINI_VERSION | tr . -`
	export GCP_SSH_TAG=$GCP_SSH_TAG
	packer build -only=googlecompute Packerfile.json
else
	echo "Unrecognized platform. Aborted."
	exit 1
fi
