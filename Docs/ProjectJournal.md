## Week Apr. 7- Apr. 13
After discussion, we have decided to implement a "clone" of Quora.com. 

#### Backend includes following GO APIs:

* User Authentication:  User SignUp, User Login, JWT token assign and verify

* User Activities:  Follow/Unfollow questions and topics, Feed questions to user in `/Home` page

* Questions: Get/Add/Update questions,answers and comments

#### Frontend:     ReactJS

### Responsibilitis:

* Yu Zhao: Frontend design and implementation
* Yueqiao Zhang: User Activities API
* Hongzhe Yang: User Authentication API
* David Ronca: Questions API

## Week Apr. 14 - Apr. 21

After discussion on Slake group and meetup after class, we have discussed following topics:

* Achitechture Design of Quora Application
* Dependency between backend GO APIs

### Progress:
* ReactJS: 
    * [API Doc](https://docs.google.com/spreadsheets/d/1M4RdDfX2pyHF5RVmjj8jFG7bgsPhhCXzO-LWUfgFXt8/edit?usp=sharing ) (for Frontend-Backend communication only, include Method, API Endpoint, request params/body schema, and response schema)
    * UI Design
    * SignUp, Login and Home Page (with faked backend)

* Quesions:

* User Activities:
* User Authentication: 
    * Set up the JWT for user 
    * Set up the MongoDB for user 
    * Able to give authentication base on tokens

### To Do List:
* Yu Zhao:
    * Work on Question Page (allow "follow question/topic, answer question functionalities) and Topic Page
    * Tune communication between Frontend and Backend
    * Test all individual microservices through frontend

* Yueqiao Zhang:

* Hongzhe Yang:
    * Match the frontend with correct request format
    * Get other JWT parts working for the other backend service

* David Ronca:
	* cwmapi (Cwoura Message API) is the message store (topics, questions, answers, and comments)
	* After some back-and-forth, we decided to use MongoDB for messages.  The primary reason was the thought that a CP database would be better for this task, and that additional scale could be achieved through sharding.  Since the top-level association is at the question, the questionId would be a natural place to shard the Cwoura message data).
	* cwmapi is written Golang, using the MGO adaptor.


## Week Apr. 22 - Apr. 28

One of road blocker we have this week is the CORS error for browser preflight request. 
The typical error message frontend received is like:
```
Access to XMLHttpRequest at ‘http://35.164.157.104:8000/msgstore/v1/topics?excludeFollowed=false’ from origin ’http://localhost:3000' has been blocked by CORS policy: Request header field authorization is not allowed by Access-Control-Allow-Headers in preflight response.
```
By import `cors` package in Go backend, now we are able to pass `Authorization` header in our request from frontend.

### Progress:

* CWMAPI (David Ronca):

	* The initial data model was not correct, as questions were tied to spaces.  We decided to drop the "Space" feature altogether, and make questions top-level.
	* Addeds support for topics as a free-form tag to questions.  That is, topics can be created freely when posting a question.
	* Added query for questions by topic.
	* Moved from Docker Mongo, to 3-node cluser.
	* Added Kong gateway.
 
* User Authorization:

* User Activity:

* Frontend (ReactJS):
	* Finished all frontend main pages (`/Home`,`/topics/{topic_name}`, `/questions/{question_id}`, `/login`,`/signup`) and all components included in thest pages. 
	* Assist Backend APIs to solve CORS error for browser preflight request
	* Start to wrap frontend up

### To Do List:

* David Ronca:

* Hongzhe Yang:

* Yu Zhao:
    * Continue the work to tune communication between Frontend and Backend
    * Test all individual microservices through frontend
    * Setup static server for Frontend and scale it up

* Yueqiao Zhang:

## Week Apr.29 - May. 4 

### Progress:

