#!/bin/bash

#echo "==============================="
#echo "INSTALLING ANSIBLE DEPENDENCIES"
#echo "-------------------------------"
apt-get update
apt-get install -y language-pack-en python-pip libffi-dev libssl-dev python-dev
sudo pip install --upgrade pip
sudo pip install markupsafe==1.1.1
sudo pip install setuptools==40.8.0
sudo pip install paramiko==1.16.0
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
