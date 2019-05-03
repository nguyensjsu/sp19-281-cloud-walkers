##Cloudwalker User Activity Mongo Shard Journal

1. Using mongo-ami to launch 1 private subnet instance for Mongo-config, 1 private subnet 
   instance for Mongo-shard1, and 1 private subnet instance for Mongo-shard2

2. For Mongo-config: 
   
   a) Log into Jumpbox

   ssh -i cmpe281-us-west-1.pem 10.0.1.47

   b) Edited Mongod config file: sudo vi /etc/mongod.conf

   change port 27017 to port 27019

   c) Update config to mongod

   sudo mongod --config /etc/mongod.conf
   sudo service mongod restart 

   d) connnect to mongo shell mongo --port 27019

   rs.initiate(

	  {
	    _id: "cfg",
	    configsvr: true,
	    members: [
	      { _id : 0, host : "10.0.1.47:27019" }
	    ]
	  }
    )
   rs.status()

3. For Mongo-shard1:
   
   a) Log into Jumpbox

   ssh -i cmpe281-us-west-1.pem 10.0.1.187

   b) Edited Mongod config file: sudo vi /etc/mongod.conf

   change port 27017 to port 27018

   c) Update config to mongod

   sudo mongod --config /etc/mongod.conf
   sudo service mongod restart 

   d) connnect to mongo shell mongo --port 27018

   rs.initiate(

	  {
	    _id: "s1",
	    members: [
	      { _id : 0, host : "10.0.1.187:27018" }
	    ]
	  }
   )
   rs.status()

4. For Mongo-shard2, repeated shard1

rs.initiate(

	  {
	    _id: "s2",
	    members: [
	      { _id : 0, host : "10.0.1.100:27018" }
	    ]
	  }
)

5. For Mongo-router, ssh -i cmpe281-us-west-1.pem 10.0.1.92

   a) in mongod.config

   uncomment replSetName

   sharding:
  	      configDB: cfg/10.0.1.47:27019

   b) sudo mongos --config /etc/mongod.conf

   c) or in jumpbox, ssh -i cmpe281-us-west-1.pem 10.0.1.92 -t "sudo mongos --transitionToAuth --fork --configdb cfg/10.0.1.47:27019 --keyFile /opt/mongodb/keyFile --logpath /var/log/mongodb.log --bind_ip 0.0.0.0"

   d) create admin user in mongos

   7) add two shards cluster in mongos

    sh.addShard("s1/10.0.1.187:27018");
	sh.addShard("s2/10.0.1.100:27018");
  
   8) set shard key 

   sh.shardCollection("cmpe281.uSpace",{_id: "hashed"}) 
