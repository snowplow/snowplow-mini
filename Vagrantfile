Vagrant.configure("2") do |config|

  config.vm.box = "ubuntu/trusty64"
  config.vm.hostname = "snowplow-mini"
  config.ssh.forward_agent = true

  # Use NFS for shared folders for better performance
  config.vm.network :private_network, ip: '192.168.50.50' # Uncomment to use NFS
  config.vm.synced_folder '.', '/vagrant', nfs: true # Uncomment to use NFS

  config.vm.network "forwarded_port", guest: 9200, host: 9200
  config.vm.network "forwarded_port", guest: 5601, host: 5601
  config.vm.network "forwarded_port", guest: 8080, host: 8080

  config.vm.provider :virtualbox do |vb|
    vb.name = Dir.pwd().split("/")[-1] + "-" + Time.now.to_f.to_i.to_s
    vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
    vb.customize [ "guestproperty", "set", :id, "--timesync-threshold", 10000 ]
    # Docker is quite hungry
    vb.memory = 2048
    vb.cpus = 4
  end

  config.vm.provision :shell do |sh|
    sh.path = "vagrant/up.bash"
  end

  # Requires Vagrant 1.7.0+
  config.push.define "publish", strategy: "local-exec" do |push|
    push.script = "vagrant/push.bash"
  end

end
