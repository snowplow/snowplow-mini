#!/bin/bash
set -e

########## FUNCTIONS #################
# Similar to Perl die
function die() {
	echo "$@" 1>&2 ; exit 1;
}

# Check if our Vagrant box is running. Expects `vagrant status` to look like:
#
# > Current machine states:
# >
# > default                   poweroff (virtualbox)
# >
# > The VM is powered off. To restart the VM, simply run `vagrant up`
#
# Parameters:
# 1. out_running (out parameter)
function is_running {
	[ "$#" -eq 1 ] || die "1 argument required, $# provided"
	local __out_running=$1

	set +e
	vagrant status | sed -n 3p | grep -q "^default\s*running (virtualbox)$"
	local retval=${?}
	set -e
	if [ ${retval} -eq "0" ] ; then
		eval ${__out_running}=1
	else
		eval ${__out_running}=0
	fi
}

# Go to parent-parent dir of this script
function cd_root() {
	source="${BASH_SOURCE[0]}"
	while [ -h "${source}" ] ; do source="$(readlink "${source}")"; done
	dir="$( cd -P "$( dirname "${source}" )/.." && pwd )"
	cd ${dir}
}
######################################

cd_root

# Precondition for running
running=0 && is_running "running"
[ ${running} -eq 1 ] || die "Vagrant guest must be running to push"

###### input validation ######
if [ "$#" -ne 2 ]
then
	echo "Inappropriate number of arguments! Aborted."
	echo "Usage: push.bash <platform> <size>"
	exit 1
else
	SP_MINI_PLATFORM="$1"
	SP_MINI_SIZE="$2"

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
			platform_cmd="export AWS_INSTANCE_TYPE=t2.large"
		elif [ "$SP_MINI_PLATFORM" == "gcp" ]
		then
			platform_cmd="export GCP_MACHINE_TYPE=n1-standard-2"
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
			platform_cmd="export AWS_INSTANCE_TYPE=t2.xlarge"
		elif [ "$SP_MINI_PLATFORM" == "gcp" ]
		then
			platform_cmd="export GCP_MACHINE_TYPE=n1-standard-4"
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
			platform_cmd="export AWS_INSTANCE_TYPE=t2.2xlarge"
		elif [ "$SP_MINI_PLATFORM" == "gcp" ]
		then
			platform_cmd="export GCP_MACHINE_TYPE=n1-standard-8"
		fi
	else
		echo "Unrecognized size! Aborted."
		echo "Available sizes; large, xlarge and xxlarge."
		exit 1
	fi
fi
###############################################

if [ "$SP_MINI_PLATFORM" == "aws" ]
then
	# Can't pass args through vagrant push so have to prompt
	read -e -p "Please enter your AWS_ACCESS_KEY_ID: " aws_access_key_id
	read -e -p "Please enter your AWS_SECRET_ACCESS_KEY: " aws_secret_access_key

	# Build AMI
	cmd="$platform_cmd && \
		export SP_MINI_SIZE=$SP_MINI_SIZE && \
		export AWS_ACCESS_KEY_ID=$aws_access_key_id && \
		export AWS_SECRET_ACCESS_KEY=$aws_secret_access_key && \
		cd /vagrant && \
		packer build -only=amazon-ebs Packerfile.json"
elif [ "$SP_MINI_PLATFORM" == "gcp" ]
then
	echo "GCP uses account.json file to authenticate."
	echo "Make sure account.json and Packerfile.json are in same directory!"
	cmd="$platform_cmd && \
		export SP_MINI_SIZE=$SP_MINI_SIZE && \
		cd /vagrant && \
		packer build -only=googlecompute Packerfile.json"
else
	echo "Unrecognized platform. Aborted."
	exit 1
fi

vagrant ssh -c "${cmd}"
