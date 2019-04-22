/*
	CWUORA API (User Dashboard) in Go 
	Uses MongoDB 
*/
package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	//"net/url"
	"encoding/json"
	"github.com/codegangsta/negroni"
	//"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    //"time"
)

// MongoDB Config
var mongodb_server = "mongodb:27017"
var mongodb_database = "cmpe281"
var mongodb_collection = "cwuora"



// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/home", homeHandler(formatter)).Methods("GET")
	mx.HandleFunc("/followspace", followedSpaceHandler(formatter)).Methods("POST")
	mx.HandleFunc("/mongoTest/{spaceid}", mongoTestHandler(formatter)).Methods("GET")

}

// API Home Handler
func homeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		 w.Header().Set("Access-Control-Allow-Origin", "*")
		/**
			Hard coded user id
		**/
		 var userId = "123456"
		//params := mux.Vars(req)
		/**
			Mongo server setup
		**/
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                fmt.Println("mongoserver panic")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        us := session.DB(mongodb_database).C("uSpace")
        uq := session.DB(mongodb_database).C("uQuestion")
        //ua := session.DB(mongodb_database).C("uAnswer")

        /**
        	Fetch TopicId from uSpace table in mongodb
        **/
        var topicResult []bson.M
		err = us.Find(bson.M{"userId": userId}).All(&topicResult)
		if err != nil {
			fmt.Println("findquery panic")
		}
        /**
         	Fetch QuestionId from uQuestion--user added questions table in mongodb
         **/
        var questionResult []bson.M
		err = uq.Find(bson.M{"userId": userId}).All(&questionResult)
		if err != nil {
			fmt.Println("findquery panic")
		}

		/**
		    Declare Response struct 
		**/ 
		var response HomeTest

 		resSpace := make([]TestTopic, len(topicResult))
 		resQuestions := make([]QuestionAPI, len(questionResult))
 		/**
 		    Space id is used for concatenate all space ids 
 		**/
		var spaceIds string
		for i := 0; i < len(topicResult); i++ {
			spaceIds += topicResult[i]["spaceId"].(string)
			fmt.Println("spaceId", spaceIds)			
		}
 		/**
 		    User added Question id is used for concatenate all uquestion ids 
 		**/
		var uquestionIds string
		for i := 0; i < len(questionResult); i++ {
			if i > 0 { 
				uquestionIds += "&"
			}
			uquestionIds += "questionId="
			uquestionIds += questionResult[i]["questionId"].(string)
			fmt.Println("uquestionIds", uquestionIds)			
		}
		/**
			Fetch Topic data from david using spaceid
		**/
		respTopic, err := http.Get("http://34.217.213.85:3000/msgstore/v1/topics")
		if err != nil {
			log.Fatalln(err)
		}

		bodyTopic, err := ioutil.ReadAll(respTopic.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var spaceTestContent []TestTopic
		json.Unmarshal(bodyTopic, &spaceTestContent)
		fmt.Println("spaceTestContent", spaceTestContent)

		for i := 0; i < len(topicResult); i++ {
			resSpace[i].Label = spaceTestContent[0].Label
		}
		response.TestTopic = resSpace

	    /**
	    	Fetch Question data from david using questionId
	    **/
	    var myQuestionUrl = "http://34.217.213.85:3000/msgstore/v1/questions?" + uquestionIds
		respQuestion, err := http.Get(myQuestionUrl)
		if err != nil {
			log.Fatalln(err)
		}

		bodyQuestion, err := ioutil.ReadAll(respQuestion.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var qTestContent []QuestionContentAPI
		json.Unmarshal(bodyQuestion, &qTestContent)
		fmt.Println("qTestContent", qTestContent)
	    /**
	    	Question Result
	    **/
		for i := 0; i < len(qTestContent); i++ {
			resQuestions[i].Id = qTestContent[i].Id
			resQuestions[i].Body = qTestContent[i].Body
			resQuestions[i].CreatedOn = qTestContent[i].CreatedOn
		}
        response.QuestionAPIs = resQuestions

	    /**
	    	Fetch Random Question data from david using questionId
	    **/

		formatter.JSON(w, http.StatusOK, response)
	}
}


// store spaces id in mongoDB 

// API Home Handler
func followedSpaceHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// // 1. connect with mongo server
		// session, err := mgo.Dial(mongodb_server)
  //       if err != nil {
  //               panic(err)
  //       }
  //       defer session.Close()
  //       session.SetMode(mgo.Monotonic, true)
  //       u := session.DB(mongodb_database).C("user")
  //       // 2. analysis client post body
  //       //fmt.Println(req.Body)
  // //       body, err := ioutil.ReadAll(req.Body)
		// // if err != nil {
		// // 	log.Fatalln(err)
		// // }

		
  //       var userP []MUserProfile
		// var userProfileapis MUserProfile
		// userProfileapis.UserId = "1234567"
		// sid := make([]string, 1)
		// quesid := make([]string, 1)
		// sid[0] = "5cb3c8ab78163fa3c9726fb3"
		// userProfileapis.Uspaces = sid
		// quesid[0] = "5cb4048478163fa3c9726fe0"
		// userProfileapis.Uquestions = quesid
		

		// // sid[0] = bson.ObjectIdHex(spaceid)
		// // quesid[0] = bson.ObjectIdHex(qid)

		// //userProfileapis.uspaces = bson.ObjectIdHex(spaceid)
		
		

		// fmt.Println("userProfileapis", userProfileapis)
		// u.Insert(userProfileapis);
		// userP = append(userP, userProfileapis)

		


		// // fmt.Println("spaceapis", spaceapis)
		// // fmt.Println("spaceapis[0] title",spaceapis[0].Title)

		// // // 3. insert into mongo server
		// // c.Insert(spaceapis[0])
		
  // //       var result bson.M
  // //       err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
  // //       if err != nil {
  // //               log.Fatal(err)
  // //       }
        
		// formatter.JSON(w, http.StatusOK, userP)
	}
}

// API MongoTest Handler
func mongoTestHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		var id string = params["spaceid"]

		// 1. connect with mongo server
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                fmt.Println("mongoserver panic")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        // 2. retrieve all data from mongodb
        //convert string to objectid
        bsonObjectID := bson.ObjectIdHex(id)
        // Query One
		var result bson.M
		err = c.Find(bson.M{"_id": bsonObjectID}).One(&result)
		if err != nil {
			fmt.Println("findquery panic")
		}

	    var title string = result["title"].(string)
	    fmt.Println("result", title)
	}
}

/****
/home

db.space.insert(
	{
		_id: 12345,
		title: "spacecontenttest",
		questions: [
			{
				_id: 123456,
				questionText: "testquestion",
				createdOn: "0001-01-01T00:00:00Z"
				answers: [
					{
						_id: 12344,
						answerText: "answertest",
						createdOn: "0001-01-01T00:00:00Z"
					}
				]
			}
		]
	}
);

db.quesiton.insert(
	{
		_id: 123456,
		questionText: "testquestion",
		createdOn: "0001-01-01T00:00:00Z"
		answers: [
			{
				_id: 12344,
				answerText: "answertest",
				createdOn: "0001-01-01T00:00:00Z"
			}
		]
	}
);

db.answer.insert(
	{
		_id: 12344,
		answerText: "answertest",
		createdOn: "0001-01-01T00:00:00Z"
	}
);

db.user.insert(
	{
		userId: "123456",
		fspaces: ["5cb3c8ab78163fa3c9726fb3"],
		fquestions: ["5cb4048478163fa3c9726fe0"]
	}
);

db.uSpace.insert(
	{
		userId: "123456",
		spaceId: "5cb3c8ab78163fa3c9726fb3"
	}
);

db.uQuestion.insert(
	{
		userId: "123456",
		questionId: "5cb4048478163fa3c9726fe0"
	}
);

db.uQuestion.insert(
	{
		userId: "123456",
		questionId: "5cb4048478163fa3c9726fdc"
	}
);

db.uSpace.insert(
	{
		userId: "123456",
		spaceId: "5cb3c8ab78163fa3c9726fb3"
		fquestions: ["5cb4048478163fa3c9726fe0"]
	}
);




**/


