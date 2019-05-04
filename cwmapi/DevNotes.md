#Cloud Walkers Project MAPI Notes
These are the notes to setup and run the Messaging API for the CWoura application.

#Riak Setup
Riak Nodes
tp-riak-n1	10.0.3.221
tp-riak-n2	10.0.1.51
tp-riak-n3	10.0.1.104

1. start riak on each node "sudo riak start"
2. build the cluster.  
	sudo riak-admin cluster join riak@10.0.3.221
3. test
	curl -i http://10.0.3.221:8098/ping
	curl -i http://10.0.1.51:8098/ping
	curl -i http://10.0.1.104:8098/ping
4. Setup classic load balancer
	 http://mt-elb-1847483819.us-west-2.elb.amazonaws.com 	
	
##Mongo Notes
Nodes
tp-mongo-c1n1	 10.0.3.17
tp-mongo-c1n2	 10.0.3.47
tp-mongo-c1n3	 10.0.3.130

// setup hosts
10.0.3.78  tp-mongo-c1n1
10.0.3.218 tp-mongo-c1n2
10.0.3.224 tp-mongo-c1n3

sudo hostnamectl set-hostname tp-mongo-c1n1
sudo hostnamectl set-hostname tp-mongo-c1n2
sudo hostnamectl set-hostname tp-mongo-c1n3

rs.initiate( {
      _id : "cmpe281",
      members: [
         { _id: 0, host: "tp-mongo-c1n1:27017" },
         { _id: 1, host: "tp-mongo-c1n2:27017" },
         { _id: 2, host: "tp-mongo-c1n3:27017" }
      ]
   })

db.createUser( {
           user: "admin",
           pwd: "query",
           roles: [{ role: "root", db: "admin" }]
   })

mongo -u admin -p query --authenticationDatabase admin
mongo -u admin -p query 10.0.3.218 --authenticationDatabase admin

### MongoID

Mongo handles UUID very poorly.  UUID is also considered a poor primary key because of random distribution of records.  Decided to use an incrementing 64-bit integer.

```js
use cmpe281
db.createCollection("counters")
db.counters.insert({_id:"spaceid",sequence_value:NumberInt(1)})
```

To add a record with incremented id

```
function getNextSequenceValue(){
    var sequenceDocument = db.counters.findAndModify({
        query:{_id: "spaces" },
        update: {$inc:{sequence_value:1}},
        new:true
    });
    return sequenceDocument.sequence_value;
}

db.spaces.insert
(
    {
        "_id" : NumberInt(getNextSequenceValue()),
        "title" : "Real Programmers use C",
        "createdOn" :  new Date(),
        "description" : "A space  for managing memory, apointers,  pointers to functions, and pointers to functions that return arrays of pointers to functions " +
        "that return int, etc.",
        "tags" : [
            {
                "tag": "Computer Programming",
            },
            {
                "tag": "Computer Science",
            },
            {
                "tag": "Software Engineering"
            }
        ]
    }
)
```

Finally decided to use Mongo objectID.

`Id      bson.ObjectId `json:"_id" bson:"_id,omitempty"``

Using OpenAPI to define RESTful API.  Online tool does not convert OpenAPI to go server, only older Swagger. 

##Generate docs and code

openapi-generator generate -g html -i ./CWMapi.yaml -o Y2G.1/

openapi-generator generate -g go-server -i ./CWMapi.yaml -o Y2G.1/

##Populate MongoDB with test data when on private VPC

1. Use local Docker mongo and Robo 3T to build the test data set.
2. Backup the data.  This command will back up all databases: `mongodump --db cmpe281 --out data/db/backup/`
3. We only care about the cmpe281 tables.  Use scp to copy ./data/backup/cmpe281 to jumbox, and then from jump box to Mongo primary node.
4. From Mongo primary, load into cluster: sudo mongorestore -u admin -p query --authenticationDatabase admin  --db cmpe281 /tmp/cmpe281/

##Setup Mongo Sharding

2 1-node shards + cfg

Following Yu's notes.
1. on each host, sudo service mongod stop
2. sudo mkdir /data
3. sudo mkdir /data/mongodb
4. sudo chown mongodb:mongodb /data/mongodb/


##Create NAT for callout to user activity

route table rtb-038eb8867f242f779