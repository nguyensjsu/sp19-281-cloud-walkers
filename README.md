# Team Cloudwalkers
# Architecture
![Architecture](https://user-images.githubusercontent.com/25470890/57170467-e6cfbb00-6dc1-11e9-88a7-799459629284.png)
# Team Members
* [David Ronca](https://github.com/)
* [Yu Zhao](https://github.com/yarns-backyard)
  - Something wrong with my github email link, all my commits from Mac are not correctly linked to my github account. You could find those commits with author "Yu Zhao" in commit history, insights/pulse or [git log](https://github.com/nguyensjsu/sp19-281-cloud-walkers/blob/master/frontend/gitlog_YuZhao.log).
  
* [Hongzhe Yang](https://github.com/)
* [Janet(Yueqiao)Zhang](https://github.com/treetree0211)

# Project Topic
The project is a "clone" of Quora(https://www.quora.com) saas platform.
![Screen Shot 2019-05-03 at 5 39 53 PM](https://user-images.githubusercontent.com/25470890/57171836-4a131a80-6dcd-11e9-9aa0-618faaed94eb.png)

## Project Github Repo Link
https://github.com/nguyensjsu/sp19-281-cloud-walkers

## Project Board Link
https://github.com/nguyensjsu/sp19-281-cloud-walkers/projects/1

## Project Journal Link
https://github.com/nguyensjsu/sp19-281-cloud-walkers/blob/master/Docs/ProjectJournal.md

## The repo has following structure:
- Docs has all project journal and reasearch resources (Contributes to All)
- Database has all User Authentication APIs go source code (Contributes to Hongzhe Yang)
- cwmapi has all Topics,Questions and Answers APIs go source code (Contributes to David Ronca)
- CWUORA-userDashboard has all User Activity APIs go source code (Contributes to Yueqiao Zhang)
- frontend has all frontend React source code (Contributes to Yu Zhao)

## Team member contributions:
1. David Ronca
- Built cwmapi, the messaging service for  CWoura.  Messages are questions, answers and comments.
- The cwmapi API is defined using OpenAPI.  API def is [here](https://app.swaggerhub.com/apis-docs/jonathannah/cwmapi/1).
- Established the data model for CWoura messages.
- SWorked on documentation and presentation

2. Yu Zhao
- Application design and [API Doc](https://docs.google.com/spreadsheets/d/1M4RdDfX2pyHF5RVmjj8jFG7bgsPhhCXzO-LWUfgFXt8/edit?usp=sharing )
- Frontend (ReactJS) implementation and test ([checklist](https://github.com/nguyensjsu/sp19-281-cloud-walkers/blob/master/Docs/Frontend/PageTest.md))
- Scale out static Frontend server on AWS and set up Kong API Gateway for all backend microservices
- Assist backend to set up sharded Mongo Cluster
- Collaborate with team members to solve CORS browser preflight error

3. Hongzhe Yang
-
-

4. Yueqiao Zhang
- Implemented User Actitives APIs including /home (Deprecated), POST /userFollow, GET /userFollow, GET /userFeed(Deprecated),
  POST /userPost in golang and using dockerfile to run the APIs 
- Implemented MongoDB two shards to store user activities data (used one private subnet instance for config, one private subnet instance for shard1, another one for shard2 and one mongos router)
- Tested with frontend and fixed issues according to frontend needs
- Contributed to JWT Token server decode and CORS issues fixed
- Contributed to project journal and other team docs
  


