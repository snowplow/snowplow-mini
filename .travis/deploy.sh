#!/bin/bash
set -e

cd $TRAVIS_BUILD_DIR

###### input validation ######
if [ "$#" -ne 1 ]
then
	echo "Inappropriate number of arguments! Aborted."
	echo "Usage: deploy.sh tag"
	exit 1
else
	export SP_MINI_PLATFORM=`echo $1 | cut -d'/' -f2`
	export SP_MINI_SIZE=`echo $1 | cut -d'/' -f3`

	if [ "$SP_MINI_PLATFORM" == "aws" -o "$SP_MINI_PLATFORM" == "gcp" ]
	then
		echo "Platform recognized!"
	else
		echo "Unrecognized platform! Aborted."
		echo "Supported platforms; aws, gcp."
		exit 1
	fi

	if [ "$SP_MINI_SIZE" == "large" ]
	then
		# prepare env vars for docker-compose
		# to be used as -Xmx jvm option for Elasticsearch & Snowplow apps
		echo -n > provisioning/roles/docker/files/.env
		echo "ES_JVM_SIZE=4g" >> provisioning/roles/docker/files/.env
		echo "SP_JVM_SIZE=512m" >> provisioning/roles/docker/files/.env
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
		# prepare env vars for docker-compose
		# to be used as -Xmx jvm option for Elasticsearch & Snowplow apps
		echo -n > provisioning/roles/docker/files/.env
		echo "ES_JVM_SIZE=8g" >> provisioning/roles/docker/files/.env
		echo "SP_JVM_SIZE=1536m" >> provisioning/roles/docker/files/.env
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
		# prepare env vars for docker-compose
		# to be used as -Xmx jvm option for Elasticsearch & Snowplow apps
		echo -n > provisioning/roles/docker/files/.env
		echo "ES_JVM_SIZE=16g" >> provisioning/roles/docker/files/.env
		echo "SP_JVM_SIZE=3g" >> provisioning/roles/docker/files/.env
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
		echo "Available sizes; large, xlarge and xxlarge."
		exit 1
	fi
fi

############## deployment ###################

if [ "$SP_MINI_PLATFORM" == "aws" ]
then
	export AWS_ACCESS_KEY_ID=$AWS_DEPLOY_ACCESS_KEY
	export AWS_SECRET_ACCESS_KEY=$AWS_DEPLOY_SECRET_KEY
	packer build -only=amazon-ebs Packerfile.json
elif [ "$SP_MINI_PLATFORM" == "gcp" ]
then
	packer build -only=googlecompute Packerfile.json
else
	echo "Unrecognized platform. Aborted."
	exit 1
fi
