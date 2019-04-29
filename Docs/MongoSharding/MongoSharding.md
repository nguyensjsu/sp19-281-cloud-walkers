## 1. Create Mongo AMI for sharded cluster
```
1. AMI:             Ubuntu Server 16.04 LTS (HVM)
2. Instance Type:   t2.micro
3. VPC:             cmpe281
4. Network:         private subnet
5. Auto Public IP:  no
6. Security Group:  mongodb-sharded-cluster 
7. SG Open Ports:   22, 27017,27018,27019
8. Key Pair:        grails-instance.pem 
```
Enable NAT gateway with public subnet in the VPC.

Install Mongo DB
```
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/4.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb.list
sudo apt update
sudo apt install mongodb-org
```

Check mongo version
```
sudo systemctl start mongod
mongod --version
```
create KeyFile
```
/*create a key file*/
openssl rand -base64 741 > keyFile
sudo mkdir -p /opt/mongodb
sudo cp keyFile /opt/mongodb

/*change owner of keyFile to mongodb and group of mongodb*/
sudo chown mongodb:mongodb /opt/mongodb/keyFile

/*change permission to user could read and write
sudo chmod 0600 /opt/mongodb/keyFile
```
Create A AMI based on above instance

launch a jump box instance to access all instance in private subnet.

## 2. Create Config Server Replica Set

For production enviorment, launch 3 instance using above mongo AMI.
```
Instance Type:   t2.micro
VPC:             cmpe281
Network:         private subnet
Auto Public IP:  no
Security Group:  mongodb-sharded-cluster 
SG Open Ports:   22, 27017,27018,27019
Key Pair:        grails-instance.pem 
```

SSH into jump box instance
```
vi lauch_config.sh
```

add following script
```
#!/bin/bash

ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.232 -t "sudo mongod --fork --logpath /var/log/mongodb.log --configsvr --replSet configset --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"
ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.82 -t "sudo mongod --fork --logpath /var/log/mongodb.log --configsvr --replSet configset --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"
ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.223 -t "sudo mongod --fork --logpath /var/log/mongodb.log --configsvr --replSet configset --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"
```
```
chmod +x launch_config.sh

./lauch_config.sh
```

SSH into any instance (e.g. 10.0.1.232)

```
ps -aux |grep mongo
```

mongod should be running in the instance

connnect to mongo shell `mongo --port 27019`:
```
rs.initiate(
{
  _id: "configset",
  configsvr:true,
  members:[
    {_id: 0, host: "10.0.1.232:27019"},
    {_id: 1, host: "10.0.1.82:27019"},
    {_id: 2, host: "10.0.1.223:27019"}
  ]
}
)

rs.status()
```

## 3. Create Sharded Replica Set

For production enviorment, launch 3 instance per shard using above mongo AMI.

```
Instance Type:   t2.micro
VPC:             cmpe281
Network:         private subnet
Auto Public IP:  no
Security Group:  mongodb-sharded-cluster 
SG Open Ports:   22, 27017,27018,27019
Key Pair:        grails-instance.pem 
```
SSH into jump box instance
```
vi lauch_shards.sh
```

add following script
```
#!/bin/bash

ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.9 -t "sudo mongod --fork --logpath /var/log/mongodb.log --shardsvr --replSet shard1 --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"
ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.233 -t "sudo mongod --fork --logpath /var/log/mongodb.log --shardsvr --replSet shard1 --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"
ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.181 -t "sudo mongod --fork --logpath /var/log/mongodb.log --shardsvr --replSet shard1 --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"

ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.119 -t "sudo mongod --fork --logpath /var/log/mongodb.log --shardsvr --replSet shard2 --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"
ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.8 -t "sudo mongod --fork --logpath /var/log/mongodb.log --shardsvr --replSet shard2 --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"
ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.43 -t "sudo mongod --fork --logpath /var/log/mongodb.log --shardsvr --replSet shard2 --keyFile /opt/mongodb/keyFile --dbpath /data/mongodb --bind_ip 0.0.0.0"
```
```
chmod +x launch_shards.sh

./lauch_shard.sh
```
SSH into any instance (e.g. 10.0.1.9)

```
ps -aux |grep mongo
```

mongod should be running in the instance

connnect to mongo shell `mongo --port 27018`:
```
rs.initiate(
...   {
...     _id : "shard1",
...     members: [
...     { _id: 0, host: "10.0.1.9:27018"},
...     { _id: 1, host: "10.0.1.233:27018"},
...     {_id:  2, host: "10.0.1.181:27018"}
...  ]
... }
... )

rs.status()
```

SSH into 10.0.1.119,connect to mongo shell `mongo --port 27018`
```
rs.initiate(
... ... {
... ... _id: "shard2",
... ... members:[
... ... {_id: 0, host: "10.0.1.119:27018"},
...     {_id: 1, host: "10.0.1.8:27018"},
...     {_id: 2, host: "10.0.1.43:27018"}
... ]
... }
... )

rs.status()
```

## 4. Connect mongos to Sharded Replica Set
launch a router instance using above mongo AMI.

```
Instance Type:   t2.micro
VPC:             cmpe281
Network:         private subnet
Auto Public IP:  no
Security Group:  mongodb-sharded-cluster 
SG Open Ports:   22, 27017,27018,27019
Key Pair:        grails-instance.pem 
```
SSH into jump box instance
```
vi lauch_router.sh
```

add following script
```
#!/bin/bash

ssh -i ~/.ssh/grails-instance.pem ubuntu@10.0.1.72 -t "sudo mongos --transitionToAuth --fork --configdb configset/10.0.1.232:27019,10.0.1.82:27019,10.0.1.223:27019 --keyFile /opt/mongodb/keyFile --logpath /var/log/mongodb.log --bind_ip 0.0.0.0"
```
```
chmod +x launch_router.sh

./lauch_router.sh
```

## 5. Add Shard to cluster
SSH into router instance, connect to mongo shell, `mongo --port 27017`

```
sh.status()
sh.addShard("shard2/10.0.1.119:27018")
sh.addShard("shard1/10.0.1.9:27018")
```

check current status
```
sh.status()
```
result:
```
--- Sharding Status ---
  sharding version: {
  	"_id" : 1,
  	"minCompatibleVersion" : 5,
  	"currentVersion" : 6,
  	"clusterId" : ObjectId("5c9ff2d63d9310d2bd9de003")
  }
  shards:
        {  "_id" : "shard1",  "host" : "shard1/10.0.1.181:27018,10.0.1.233:27018,10.0.1.9:27018",  "state" : 1 }
        {  "_id" : "shard2",  "host" : "shard2/10.0.1.119:27018,10.0.1.43:27018,10.0.1.8:27018",  "state" : 1 }
  active mongoses:
        "4.0.8" : 1
  autosplit:
        Currently enabled: yes
  balancer:
        Currently enabled:  yes
        Currently running:  no
        Failed balancer rounds in last 5 attempts:  0
        Migration Results for the last 24 hours:
                No recent migrations
  databases:
        {  "_id" : "config",  "primary" : "config",  "partitioned" : true }

```
# Test

connect to mongos shell
```
use sharding-test       //create a new database
sh.enableSharding("sharding-test")

sh.shardCollection("sharding-test.bios",{_id: "hashed"})        //choose "_id" as shard key

```

Using bios collection from previous lab, insert collection into db
check current shard distributions
```
db.bios.getShardDistribution()
```

results:

```
Shard shard1 at shard1/10.0.1.181:27018,10.0.1.233:27018,10.0.1.9:27018
 data : 1KiB docs : 5 chunks : 2
 estimated data per chunk : 705B
 estimated docs per chunk : 2

Shard shard2 at shard2/10.0.1.119:27018,10.0.1.43:27018,10.0.1.8:27018
 data : 2KiB docs : 5 chunks : 2
 estimated data per chunk : 1KiB
 estimated docs per chunk : 2

Totals
 data : 3KiB docs : 10 chunks : 4
 Shard shard1 contains 39.96% data, 50% docs in cluster, avg obj size on shard : 282B
 Shard shard2 contains 60.03% data, 50% docs in cluster, avg obj size on shard : 423B
```

Bios Collection is sharded into different shards (shard1 and shard2) with `"_id"` and shared key