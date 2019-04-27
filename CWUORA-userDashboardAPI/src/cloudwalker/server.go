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
    "github.com/dgrijalva/jwt-go"
    //"time"
    //"math/rand"
    "strings"
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
	mx.HandleFunc("/userFollow", followHandler(formatter)).Methods("POST")
	mx.HandleFunc("/userPost", userPostHandler(formatter)).Methods("POST")
	mx.HandleFunc("/userFollow", followListHandler(formatter)).Methods("GET")
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
        	Fetch Topic Labels from uSpace table in mongodb
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
 		ranQuestions := make([]QuestionAPI, 5)
 		/**
 		    Space id is used for concatenate all space ids 
 		**/
		//var spaceIds string
		for i := 0; i < len(topicResult); i++ {
			resSpace[i].Label = topicResult[i]["spaceId"].(string)		
		}
		response.TestTopic = resSpace
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
		randomQuestion, err := http.Get("http://34.217.213.85:3000/msgstore/v1/questions")
		if err != nil {
			log.Fatalln(err)
		}

		bodyRanQuestion, err := ioutil.ReadAll(randomQuestion.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var qRanContent []QuestionContentAPI
		json.Unmarshal(bodyRanQuestion, &qRanContent)
		fmt.Println("qRanContent", qRanContent)
	    /**
	    	Question Result
	    **/
		for i := 0; i < 5; i++ {
			ranQuestions[i].Id = qRanContent[i].Id
			ranQuestions[i].Body = qRanContent[i].Body
			ranQuestions[i].CreatedOn = qRanContent[i].CreatedOn
		}
        response.RandomAPIs = ranQuestions


		formatter.JSON(w, http.StatusOK, response)
	}
}


// API Follow/Unfollow Handler POST
func followHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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
        uq := session.DB(mongodb_database).C("uFQuestion")
		/**
			Get Post body
		**/        
        body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(body)

		var followResult PostFollow
		json.Unmarshal(body, &followResult)

		/** Get user ID from JWT, header
		**/
		tokenStrWithSpace := req.Header.Get("Authorization")
		//var tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTYzNDYxODEsImlkIjoiNWNjMDAwYTk3MmM5YmZmZjEwNzU4MWUxIn0.r_T2oKqsmK6PjHZ-lZQROD3u1gAOd3uxjRwLrk8LanQ"
		tokenStr := strings.Split(string(tokenStrWithSpace), " ")[1]
		//var tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTYzNDYxODEsImlkIjoiNWNjMDAwYTk3MmM5YmZmZjEwNzU4MWUxIn0.r_T2oKqsmK6PjHZ-lZQROD3u1gAOd3uxjRwLrk8LanQ"
		
        hmacSecretString := "secret"// Value
        hmacSecret := []byte(hmacSecretString)
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
             // check token signing method etc
             return hmacSecret, nil
        })
        claims, ok := token.Claims.(jwt.MapClaims)
        // ; ok && token.Valid {
        //     return claims, true
        // } else {
        //     log.Printf("Invalid JWT Token")
        //     return nil, false
        // }

        fmt.Println("claims", claims)
        userId := claims["id"].(string)

        fmt.Println("decodedUserid", userId)
        fmt.Println("ok", ok)

		//var userId = "888888"
		var action = followResult.Action
		

		if action == "topic" {			
			//topic.Uspaces = followId 
			var followId = followResult.Id
			var ifFollow = followResult.Unfollow
			
			for i := 0; i < len(followId); i++ {
				var topic MUserSpace
				topic.UserId = userId
				topic.Uspaces = followId[i]
				if ifFollow == false {
					us.Insert(topic)
				} else {
					us.Remove(topic)
				}
			}
		}

		if action == "question" {
			var followId = followResult.Id
			var ifFollow = followResult.Unfollow

			var question MUserFQuestion
			question.UserId = userId
			question.FollowedQ = followId[0] 
			if ifFollow == false {
				uq.Insert(question)
			} else {
				uq.Remove(question)
			}
			
		}

		var response Success
		response.Success = true

		formatter.JSON(w, http.StatusOK, response)
	}
}


// API Post content Handler
func userPostHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		/**
			Mongo server setup
		**/
		w.Header().Set("Access-Control-Allow-Origin", "*")
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                fmt.Println("mongoserver panic")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        uq := session.DB(mongodb_database).C("uQuestion")
        ua := session.DB(mongodb_database).C("uAnswer")
		/**
			Get Post body
		**/        
        body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(body)

		var postResult PostContent
		json.Unmarshal(body, &postResult)

		/**
			Hard code userid for testing
		**/   
		//var userId = "888888"
		tokenStrWithSpace := req.Header.Get("Authorization")
		//var tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTYzNDYxODEsImlkIjoiNWNjMDAwYTk3MmM5YmZmZjEwNzU4MWUxIn0.r_T2oKqsmK6PjHZ-lZQROD3u1gAOd3uxjRwLrk8LanQ"
		tokenStr := strings.Split(string(tokenStrWithSpace), " ")[1]
        hmacSecretString := "secret"// Value
        hmacSecret := []byte(hmacSecretString)
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
             // check token signing method etc
             return hmacSecret, nil
        })
        claims, ok := token.Claims.(jwt.MapClaims)
        // ; ok && token.Valid {
        //     return claims, true
        // } else {
        //     log.Printf("Invalid JWT Token")
        //     return nil, false
        // }

        fmt.Println("claims", claims)
        userId := claims["id"].(string)

        fmt.Println("decodedUserid", userId)
        fmt.Println("ok", ok)
		var action = postResult.Action
		var postId = postResult.Id

		if action == "question" {
			var question MUserQuestion
			question.UserId = userId
			question.Uquestions = postId 
			uq.Insert(question)
		}

		if action == "answer" {
			var answer MUserAnswer
			answer.UserId = userId
			answer.UAnswers = postId 
			ua.Insert(answer)
		}
		
		var response Success
		response.Success = true
        
		formatter.JSON(w, http.StatusOK, response)
	}
}

// API User Follow GET content Handler
func followListHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		/**
			Mongo server setup
		**/
		w.Header().Set("Access-Control-Allow-Origin", "*")
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                fmt.Println("mongoserver panic")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        us := session.DB(mongodb_database).C("uSpace")
        //ua := session.DB(mongodb_database).C("uAnswer")
		/**
			Get Post body
		// **/        
  //       body, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// fmt.Println(body)

		// var postResult PostContent
		// json.Unmarshal(body, &postResult)

		/**
			Hard code userid for testing
		**/   
		//var userId = "888888"
		tokenStrWithSpace := req.Header.Get("Authorization")
		//var tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTYzNDYxODEsImlkIjoiNWNjMDAwYTk3MmM5YmZmZjEwNzU4MWUxIn0.r_T2oKqsmK6PjHZ-lZQROD3u1gAOd3uxjRwLrk8LanQ"
		tokenStr := strings.Split(string(tokenStrWithSpace), " ")[1]
        hmacSecretString := "secret"// Value
        hmacSecret := []byte(hmacSecretString)
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
             // check token signing method etc
             return hmacSecret, nil
        })
        claims, ok := token.Claims.(jwt.MapClaims)
        // ; ok && token.Valid {
        //     return claims, true
        // } else {
        //     log.Printf("Invalid JWT Token")
        //     return nil, false
        // }

        fmt.Println("claims", claims)
        userId := claims["id"].(string)

        fmt.Println("decodedUserid", userId)
        fmt.Println("ok", ok)

        /** Get user followed space ID from Mongo
        **/
        var spaceResult []bson.M
		err = us.Find(bson.M{"userId": userId}).All(&spaceResult)
		if err != nil {
			fmt.Println("findquery panic")
		}

        var response FollowSpaceList
		resSpace := make([]TestTopic, len(spaceResult))

		for i := 0; i < len(spaceResult); i++ {
			resSpace[i].Label = spaceResult[i]["spaceId"].(string)		
		}
		response.FollowSpace = resSpace
        
		formatter.JSON(w, http.StatusOK, response)
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
		userId: "5cc000a972c9bfff107581e1",
		spaceId: "Tourism"
	}
);
db.uSpace.insert(
	{
		userId: "5cc000a972c9bfff107581e1",
		spaceId: "Student"
	}
);

db.uQuestion.insert(
	{
		userId: "123456",
		questionId: "5cb4048478163fa3c9726fe0"
	}
);

db.uFQuestion.insert(
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


