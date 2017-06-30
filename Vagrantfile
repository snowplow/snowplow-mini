Vagrant.configure("2") do |config|

  config.vm.box = "ubuntu/trusty64"
  config.vm.hostname = "snowplow-mini"
  config.ssh.forward_agent = true
  # Set the name of the VM. See: http://stackoverflow.com/a/17864388/100134
  #config.vm.define :snowplowmini do |snowplowmini|
  #end

  # Use NFS for shared folders for better performance
  config.vm.network :private_network, ip: '192.168.50.50' # Uncomment to use NFS
  config.vm.synced_folder '.', '/vagrant', nfs: true # Uncomment to use NFS

  config.vm.network "forwarded_port", guest: 2000, host: 2000
  config.vm.network "forwarded_port", guest: 3000, host: 3000
  config.vm.network "forwarded_port", guest: 8080, host: 8080
  config.vm.network "forwarded_port", guest: 9200, host: 9200
  config.vm.network "forwarded_port", guest: 5601, host: 5601
  config.vm.network "forwarded_port", guest: 8081, host: 8081
  config.vm.network "forwarded_port", guest: 10000, host: 10000

  config.vm.provider :virtualbox do |vb|
    vb.name = Dir.pwd().split("/")[-1] + "-" + Time.now.to_f.to_i.to_s
    vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
    vb.customize [ "guestproperty", "set", :id, "--timesync-threshold", 10000 ]
    vb.memory = 4096 
    vb.cpus = 1
  end

  config.vm.provision :shell do |sh|
    sh.path = "vagrant/up.bash"
  end

  # Requires Vagrant 1.7.0+
  config.push.define "publish", strategy: "local-exec" do |push|
    push.script = "vagrant/push.bash"
  end

end
