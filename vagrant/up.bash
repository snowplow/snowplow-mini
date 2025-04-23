#!/bin/bash

#echo "==============================="
#echo "INSTALLING ANSIBLE DEPENDENCIES"
#echo "-------------------------------"
export DEBIAN_FRONTEND=noninteractive
apt update && apt upgrade -y
sudo sh -c 'echo "ubuntu ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers' # see https://askubuntu.com/questions/192050/how-to-run-sudo-command-with-no-password

#echo "=================="
#echo "INSTALLING ANSIBLE"
#echo "------------------"
sudo apt install software-properties-common -y
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible=11.5.0-1ppa~noble -y

#echo "=========================================="
#echo "RUNNING PLAYBOOKS WITH ANSIBLE*"
#echo "------------------------------------------"

vagrant_dir=/vagrant/vagrant
cd $vagrant_dir/..
ansible-playbook -i provisioning/inventory provisioning/local_setup.yml --connection=local --become
