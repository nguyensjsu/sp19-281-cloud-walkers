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
    "time"
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
		//params := mux.Vars(req)
		//var spaceid string = "5cb3c8ab78163fa3c9726fb3"
		/*************Mongo server setup*******************/
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                fmt.Println("mongoserver panic")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        ms := session.DB(mongodb_database).C("space")
        mq := session.DB(mongodb_database).C("question")
        ma := session.DB(mongodb_database).C("answer")

		/************** store all info, fetch data from david and store in mongodb ******/
		resp, err := http.Get("http://34.217.213.85:3000/msgstore/v1/spaces?depth=2")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var spaceAllContent []SpaceContentAPI
		json.Unmarshal(body, &spaceAllContent)
		/**************** Convert data to mongo ****************************************/
		/* create mongo space struct */
		mspaces := []*MSpace{}
		mspace := new(MSpace)

		for i := 0; i < len(spaceAllContent); i++ {

			mspace = new(MSpace)
			mspace.Id = spaceAllContent[i].Id
			mspace.Title = spaceAllContent[i].Title

			/******** create mongo question struct *********/
			var count = len(spaceAllContent[i].Questions)
			mquestions := make([]MQuestion, count)

			var qcontent []QuestionContentAPI
			qcontent = spaceAllContent[i].Questions

			for j := 0; j < count; j++ {
				mquestions[j].Id = qcontent[j].Id
				mquestions[j].Body = qcontent[j].Body
				mquestions[j].CreatedOn = qcontent[j].CreatedOn

				/******** create mongo answer struct *********/
				var acount = len(qcontent[j].Answers)
				manswers := make([]MAnswer, acount)

				var acontent []AnswerContentAPI
				acontent = qcontent[j].Answers

				for k := 0; k < acount; k++ {
					manswers[k].Id = acontent[k].Id
					manswers[k].Body = acontent[k].Body
					manswers[k].CreatedOn = acontent[k].CreatedOn
					/*********insert answer data into mongo*******************/
					ma.Insert(manswers[k])
				}

				mquestions[j].Answers = manswers
				/*********insert question data into mongo*******************/
				mq.Insert(mquestions[j])
			}

			mspace.Questions = mquestions
			/*********insert space data into mongo*******************/
			ms.Insert(mspace)
			mspaces = append(mspaces, mspace)
		}

		/**************** Hard code userID for test ******************************/
		var userId = "1234567"
		/****************** Get user space & question id *******************************/
		var uSpacesId []interface{}
		var uQuestionId []interface{}
		u := session.DB(mongodb_database).C("user")
		var userResult bson.M
		err = u.Find(bson.M{"userId": userId}).One(&userResult)
		if err != nil {
			fmt.Println("findquery panic")
		}

		uSpacesId = userResult["fspaces"].([]interface{})
		uQuestionId = userResult["fquestions"].([]interface{})

        // // find id from mongodb, using space id to find question id
        // //convert space id to mongo id
        // 

  //       var spaceResult bson.M
		// err = c.Find(bson.M{"_id": bsonObjectID}).One(&spaceResult)
		// if err != nil {
		// 	fmt.Println("findquery panic")
		// }

	 //    var selectQuestions []interface{}
	 //    selectQuestions = spaceResult["questions"].([]interface{})
	 //    var qsCount int = len(selectQuestions)
	 //    fmt.Println(qsCount)
	 //    // add question to frontend

	 //    var result bson.M
	 //    result = selectQuestions[0].(bson.M)
	  
	 //    // convert interface to array
	 //    fmt.Println("result", result["_id"])

	 //    // // generate frontend json
	     var response Home
	 //    //var resSpace []SpaceAPI
	     resSpace := make([]SpaceAPI, len(uSpacesId))
	     resQuestions := make([]QuestionAPI, len(uQuestionId))

	     /********** Get all followed spaces **********************/
	     for i := 0; i < len(uSpacesId); i++ {
	     	var spaceResult bson.M
	     	var fspaceId = uSpacesId[i].(string)
	     	bsonObjectID := bson.ObjectIdHex(fspaceId)

	     	s := session.DB(mongodb_database).C("space")
	     	err = s.Find(bson.M{"_id": bsonObjectID}).One(&spaceResult)
			if err != nil {
				fmt.Println("findquery panic")
			}

			resSpace[i].Id = spaceResult["_id"].(bson.ObjectId)
	        resSpace[i].Title = spaceResult["title"].(string)
	     }

	     /********** Get all followed spaces **********************/
	     for i := 0; i < len(uQuestionId); i++ {
	     	var spaceResult bson.M
	     	var qspaceId = uQuestionId[i].(string)
	     	bsonObjectID := bson.ObjectIdHex(qspaceId)

	     	q := session.DB(mongodb_database).C("question")
	     	err = q.Find(bson.M{"_id": bsonObjectID}).One(&spaceResult)
			if err != nil {
				fmt.Println("findquery panic")
			}

			resQuestions[i].Id = spaceResult["_id"].(bson.ObjectId)
	        resQuestions[i].Body = spaceResult["body"].(string)
	        resQuestions[i].CreatedOn = spaceResult["createdon"].(time.Time)
	     }
	 //    // add content to resSpace
	 //    resSpace[0].Id = spaceResult["_id"].(bson.ObjectId)
	 //    resSpace[0].Title = spaceResult["title"].(string)

	 //    // add content to resQuestions
	 //    for i := 0; i < qsCount; i++ {
	 //    	var r bson.M
	 //    	r = selectQuestions[i].(bson.M)
	 //    	resQuestions[i].Id = r["_id"].(bson.ObjectId)
	 //    }

	       response.SpaceAPIs = resSpace
	       response.QuestionAPIs = resQuestions

				
		formatter.JSON(w, http.StatusOK, response)
	}
}


// store spaces id in mongoDB 

// API Home Handler
func followedSpaceHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 1. connect with mongo server
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        u := session.DB(mongodb_database).C("user")
        // 2. analysis client post body
        //fmt.Println(req.Body)
  //       body, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		
        var userP []MUserProfile
		var userProfileapis MUserProfile
		userProfileapis.UserId = "1234567"
		sid := make([]string, 1)
		quesid := make([]string, 1)
		sid[0] = "5cb3c8ab78163fa3c9726fb3"
		userProfileapis.Uspaces = sid
		quesid[0] = "5cb4048478163fa3c9726fe0"
		userProfileapis.Uquestions = quesid
		

		// sid[0] = bson.ObjectIdHex(spaceid)
		// quesid[0] = bson.ObjectIdHex(qid)

		//userProfileapis.uspaces = bson.ObjectIdHex(spaceid)
		
		

		fmt.Println("userProfileapis", userProfileapis)
		u.Insert(userProfileapis);
		userP = append(userP, userProfileapis)

		


		// fmt.Println("spaceapis", spaceapis)
		// fmt.Println("spaceapis[0] title",spaceapis[0].Title)

		// // 3. insert into mongo server
		// c.Insert(spaceapis[0])
		
  //       var result bson.M
  //       err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
  //       if err != nil {
  //               log.Fatal(err)
  //       }
        
		formatter.JSON(w, http.StatusOK, userP)
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




**/


