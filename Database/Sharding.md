#### Basic set up 
Following the MongoDB set up, and the architecture I have above, I need 10 instance in total. There are 9 of them that will sit within my private subnet, and I will connect to them through my JumpBox instance. I will then config two of them as the config server for those two shards that I will have, and two replica set with 3 instances in each of them, and a router instance of mongos. 

1. Launch 9 instances with the mongo-ami that I had earlier. named them as (Here I used linux AMI instead of ubuntu)
#
	Shard-1-a 10.0.21.159
	Shard-1-b 10.0.21.25
	Shard-1-c 10.0.21.176
	Shard-2-a 10.0.21.243
	Shard-2-b 10.0.21.89
	Shard-2-c10.0.21.113
	Config-server-1 10.0.21.166
	Config-server-2 10.0.21.110
	mongos 10.0.21.160

2. Configure two servers. 
# 
	sudo mkdir -p /data/db 
	sudo chown -R mongod:mongod /data/db 
	
	mongod --configsvr --replSet <replica set name> --dbpath /data/db	--port 27019 --logpath /var/log/mongodb/mongod.log 
	ps -aux | grep mongod (check if it is running)
	sudo kill -9 <process ID>

I tried to make the command more readable by change the config file. 
#			
	sudo nano /etc/mongod.conf 
	In the file: change # where and how to store data/ 
				dbPath: /var/lib/mongo.   -> /data/db
			      # Network interfaces
				port: 27017 -> 27019
				comment out bindIP: 127.0.0.1 line 
			      replication: 
				replSetName: cmpe
			      sharing: 
				clusterRole: configsvr
	sudo mongod --config /etc/mongod.conf --logpath /var/log/mongodb/mongod.log
	Check: ps -aux | grep mongod 
	

Now connect the mongo shell to one of the config server members by specifying the port number. 
#	
	mongo -port 27019
	> rs.initiate (
	{
		_id: "cmpe",     
		configsvr: true, 
		members: [
			{ _id : 0, host : "10.0.21.166:27019"},
			{ _id : 1, host: "10.0.21.110:27019"},			
		]
	}
)
Check status: rs.status()

3. Chnage config file as well for shared replica set on all 6 instances. 
#
	Shard-1-a 10.0.21.159
	Shard-1-b 10.0.21.25
	Shard-1-c 10.0.21.176
	Shard-2-a 10.0.21.243
	Shard-2-b 10.0.21.89
	Shard-2-c10.0.21.113

	sudo mkdir -p /data/db  
	sudo chown -R mongod:mongod /data/db
	sudo nano /etc/mongod.conf 
	In the file: change # where and how to store data/ 
				dbPath: /var/lib/mongo.   -> /data/db
			      # Network interfaces
				port: 27017 -> 27018
				comment out bindIP: 127.0.0.1 line 
			      replication: 
				replSetName: <replica set name, rs0, rs1, .. > 
			      sharing: 
				clusterRole: shardsvr
After the configuration is done, start the mongoldb service on all instances. 
# 
	sudo mongod --config /etc/mongod.conf --logpath /var/log/mongodb/mongod.log 				

4. Join the two shard clusters 

#	
	mongo -port 27018
	> rs.initiate (
	{
		_id: "rs0",     
		configsvr: true, 
		members: [
			{ _id : 0, host : "10.0.21.159:27019"},
			{ _id : 1, host: "10.0.21.25:27019"},			
			{ _id : 2, host: "10.0.21.76:27019"}
		]
	}
)
#	
	rs.initiate (
		{
			_id: "rs1",     
			members: [
				{ _id : 0, host : "10.0.21.243:27018"},
				{ _id : 1, host: "10.0.21.89:27018"}, 
				{ _id : 2, host : "10.0.21.113:27018"}
			]
		}
	)


5. Configure the Mongos query router. 
#
	Need Additional package: 
	sudo yum install -y mongodb-org.mongos 
	On command line, run mongos: 
	mongos --configdb cmpe/<ip1>:27019, <ip2>:27019
	(replica set name, and ip of the replica) -> can do this in the configuration file 
	sudo nano /etc/mongod.conf 
	Comment out: #storage: 
				 #dbPath: /var/lib/mongo
				 #journal: 
				 #	enabled: true 
				
				#bindIP: ...  
	sharding: 
			configDB: cmpe/<ip1>:27019, <ip2>:27019

Start mongo: 
	sudo mongos --config /etc/mongod.conf --fork --logpath /var/log/mongodb/mongod.log