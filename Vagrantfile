Vagrant.configure("2") do |config|

  config.vm.box = "ubuntu/jammy64"
  config.vm.hostname = "snowplow-mini"
  config.ssh.forward_agent = true

  # Use NFS for shared folders for better performance
  config.vm.network :private_network, ip: '192.168.56.56' # Uncomment to use NFS
  if Vagrant::Util::Platform.windows? then
    config.vm.synced_folder '.', '/vagrant' # NFS not supported on Windows
  else
    config.vm.synced_folder '.', '/vagrant', nfs: true # Uncomment to use NFS
  end

  config.vm.network "forwarded_port", guest: 80, host: 2000     # Caddy insecure
  config.vm.network "forwarded_port", guest: 8443, host: 2443   # Caddy secure
  config.vm.network "forwarded_port", guest: 4171, host: 4171   # nsqadmin
  config.vm.network "forwarded_port", guest: 8080, host: 8080   # collector
  config.vm.network "forwarded_port", guest: 8093, host: 8093   # metrics
  config.vm.network "forwarded_port", guest: 9200, host: 9200   # elasticsearch
  config.vm.network "forwarded_port", guest: 5601, host: 5601   # kibana
  config.vm.network "forwarded_port", guest: 8081, host: 8081   # iglu-server
  config.vm.network "forwarded_port", guest: 10000, host: 10000 # control-plane

  config.vm.provider :virtualbox do |vb|
    vb.name = Dir.pwd().split("/")[-1] + "-" + Time.now.to_f.to_i.to_s
    vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
    vb.customize [ "guestproperty", "set", :id, "--timesync-threshold", 10000 ]
    vb.memory = 8192
    vb.cpus = 2
  end

  config.vm.provision :shell do |sh|
    sh.path = "vagrant/up.bash"
  end
end
