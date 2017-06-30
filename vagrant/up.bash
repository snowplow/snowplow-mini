#!/bin/bash
set -e

vagrant_dir=/vagrant/vagrant
bashrc=/home/vagrant/.bashrc

echo "========================================"
echo "INSTALLING PERU AND ANSIBLE DEPENDENCIES"
echo "----------------------------------------"
apt-get update
apt-get install -y language-pack-en git unzip libyaml-dev libssl-dev python3-pip python-pip python3-dev python-yaml python-paramiko python-jinja2
export LC_ALL=C  #see https://stackoverflow.com/a/37112094
sudo sh -c 'echo "ubuntu ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers' # see https://askubuntu.com/questions/192050/how-to-run-sudo-command-with-no-password

echo "==============="
echo "INSTALLING PERU"
echo "---------------"
sudo pip3 install peru
sudo pip install setuptools
sudo pip install ansible

echo "================================"
echo "CLONING AND PLAYBOOKS WITH PERU"
echo "--------------------------------"
cd ${vagrant_dir} && peru sync -v
echo "... done"

cp -r oss-playbooks/roles/base ../provisioning/roles/.
cp -r oss-playbooks/roles/packer ../provisioning/roles/.
cp -r oss-playbooks/roles/nodejs ../provisioning/roles/. 
cp -r oss-playbooks/roles/typescript ../provisioning/roles/.

echo "=========================================="
echo "RUNNING PLAYBOOKS WITH ANSIBLE*"
echo "------------------------------------------"

cd ..
ansible-playbook -e provision=vagrant -i provisioning/inventory provisioning/vagrant-playbook.yml --connection=local --sudo

