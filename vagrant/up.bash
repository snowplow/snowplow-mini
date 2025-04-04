#!/bin/bash

#echo "==============================="
#echo "INSTALLING ANSIBLE DEPENDENCIES"
#echo "-------------------------------"
export DEBIAN_FRONTEND=noninteractive
apt-get update
apt-get install -y language-pack-en libffi-dev libssl-dev python3-dev python3-pip
sudo pip install --upgrade pip
sudo pip install --upgrade cryptography
sudo sh -c 'echo "ubuntu ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers' # see https://askubuntu.com/questions/192050/how-to-run-sudo-command-with-no-password

#echo "=================="
#echo "INSTALLING ANSIBLE"
#echo "------------------"
sudo pip install ansible==2.8.1

#echo "=========================================="
#echo "RUNNING PLAYBOOKS WITH ANSIBLE*"
#echo "------------------------------------------"

vagrant_dir=/vagrant/vagrant
cd $vagrant_dir/..
ansible-playbook -i provisioning/inventory provisioning/local_setup.yml --connection=local --become
