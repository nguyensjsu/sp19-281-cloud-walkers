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
	mx.HandleFunc("/home/{spaceid}", homeHandler(formatter)).Methods("GET")
	mx.HandleFunc("/followspace/{userid}", followedSpaceHandler(formatter)).Methods("POST")
	mx.HandleFunc("/mongoTest/{spaceid}", mongoTestHandler(formatter)).Methods("GET")

}

// API Home Handler
func homeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		var spaceid string = params["spaceid"]

		// 1. make request to get space content, send json back
		//s := "34.217.213.85:3000/msgstore/v1/spaces?depth=0"

		resp, err := http.Get("http://34.217.213.85:3000/msgstore/v1/spaces?depth=1")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		/********************************************************/
		var depth1 []Depth1
		json.Unmarshal(body, &depth1)
		//fmt.Println(depth1)
		// get homepage followed spaces
		spaces := []*MongoInfoSpace{}
		space := new(MongoInfoSpace)
		

		for i := 0; i < len(depth1); i++ {
			var id = depth1[i].Id
			var title = depth1[i].Title
			space = new(MongoInfoSpace)
			space.Id = id
			space.Title = title
			fmt.Println("depth1", i)
			fmt.Println("title", title)
			//fmt.Println("title", depth1[i].Questions)
			// create question list for space
			var count = len(depth1[i].Questions)
			//questions := [count]SpaceQuestions{}
			fmt.Println("count", count)
			questions := make([]SpaceQuestions, count)
			//sQuestions := new(SpaceQuestions)
			for j := 0; j < count; j++ {
				questions[j].Id = depth1[i].Questions[j].Id
			}
			space.Questions = questions
			spaces = append(spaces, space)
		}


		// insert content to mongo
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                fmt.Println("mongoserver panic")
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        // insert space 
        c := session.DB(mongodb_database).C("space")
        for i := 0; i < len(spaces); i++ {
         	c.Insert(spaces[i])
        }
        // find id from mongodb, using space id to find question id
        //convert space id to mongo id
        bsonObjectID := bson.ObjectIdHex(spaceid)

        var spaceResult bson.M
		err = c.Find(bson.M{"_id": bsonObjectID}).One(&spaceResult)
		if err != nil {
			fmt.Println("findquery panic")
		}

	    var selectQuestions []interface{}
	    selectQuestions = spaceResult["questions"].([]interface{})
	    var qsCount int = len(selectQuestions)
	    fmt.Println(qsCount)
	    // add question to frontend

	    var result bson.M
	    result = selectQuestions[0].(bson.M)
	  
	    // convert interface to array
	    fmt.Println("result", result["_id"])

	    // // generate frontend json
	    var response Home
	    //var resSpace []SpaceAPI
	    resSpace := make([]SpaceAPI, 1)
	    resQuestions := make([]QuestionAPI, qsCount)
	    // add content to resSpace
	    resSpace[0].Id = spaceResult["_id"].(bson.ObjectId)
	    resSpace[0].Title = spaceResult["title"].(string)

	    // add content to resQuestions
	    for i := 0; i < qsCount; i++ {
	    	var r bson.M
	    	r = selectQuestions[i].(bson.M)
	    	resQuestions[i].Id = r["_id"].(bson.ObjectId)
	    }

	    response.SpaceAPIs = resSpace
	    response.QuestionAPIs = resQuestions

				
		formatter.JSON(w, http.StatusOK, response)
		//formatter.JSON(w, http.StatusOK, spaceResult)
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
        c := session.DB(mongodb_database).C(mongodb_collection)
        // 2. analysis client post body
        //fmt.Println(req.Body)
        body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("body",body)

		var spaceapis []SpaceAPI
		json.Unmarshal(body, &spaceapis)

		fmt.Println("spaceapis", spaceapis)
		fmt.Println("spaceapis[0] title",spaceapis[0].Title)

		// 3. insert into mongo server
		c.Insert(spaceapis[0])
		
  //       var result bson.M
  //       err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
  //       if err != nil {
  //               log.Fatal(err)
  //       }
        
		// formatter.JSON(w, http.StatusOK, result)
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
			{_id: 123456},
			{_id: 234566} ]
	}
);

db.quesiton.insert(
	{
		_id: 23456,
		body: "what is cs",
		createdOn: "0001-01-01T00:00:00Z"
	}
);





**/


