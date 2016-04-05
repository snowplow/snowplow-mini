#!/bin/bash
set -e

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

cd_root

# Precondition for running
running=0 && is_running "running"
[ ${running} -eq 1 ] || die "Vagrant guest must be running to push"

# Can't pass args thru vagrant push so have to prompt
read -e -p "Please enter your AWS_ACCESS_KEY_ID: " aws_access_key_id
read -e -p "Please enter your AWS_SECRET_ACCESS_KEY: " aws_secret_access_key

# Build AMI
cmd="export AWS_ACCESS_KEY_ID=$aws_access_key_id && \
     export AWS_SECRET_ACCESS_KEY=$aws_secret_access_key && \
     cd /vagrant/ui && npm install && tsc -p js --outDir dist/ && browserify dist/SnowplowMiniApp.js -o dist/bundle.js && \
     cd /vagrant && \
     packer build Packerfile.json"
vagrant ssh -c "${cmd}"

exit 0
