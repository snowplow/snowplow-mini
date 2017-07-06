#!/bin/bash

#echo "==============================="
#echo "INSTALLING ANSIBLE DEPENDENCIES"
#echo "-------------------------------"
apt-get update
apt-get install -y language-pack-en python-pip python-paramiko
sudo sh -c 'echo "ubuntu ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers' # see https://askubuntu.com/questions/192050/how-to-run-sudo-command-with-no-password

#echo "=================="
#echo "INSTALLING ANSIBLE"
#echo "------------------"
sudo pip install setuptools
sudo pip install ansible

#echo "=========================================="
#echo "RUNNING PLAYBOOKS WITH ANSIBLE*"
#echo "------------------------------------------"

vagrant_dir=/vagrant/vagrant
cd $vagrant_dir/..
ansible-playbook -i provisioning/inventory provisioning/with_building_ui.yml --connection=local --sudo

