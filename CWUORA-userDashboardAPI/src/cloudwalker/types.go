package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type userProfile struct {
	userId          	string    `json:"userId" bson:"userId"`	
	followedspaces		[]Space    `json:"followedspaces" bson:"followedspaces"`
	followedquestions   []Question `json:"followedquestions" bson:"followedquestions"`
}

type Space struct {
	SpaceId         string     `json:"spaceId"` 
	Title			string	`json:"title"`
	CreatedOn      time.Time `json:"createdOn"`
	Description		string  `json:"description"`
}

type Question struct {
	QuestionId      string     `json:"questionId"` 
	SpaceId         string     `json:"spaceId"` 
	Body			string     `json:"body"` 
	CreatedOn		time.Time  `json:"createdOn"`
	CreatedBy		string    `json:"createdBy"` 
	Answers			[]Answer  `json:"answers"` 
}

type Answer struct {
	AnswerId	    string   `json:"answerId"` 
	Content			string   `json:"content"` 
}

type SpaceAPI struct {
	Id 				bson.ObjectId 	`json:"_id" bson:"_id"`
	//SpaceId         string	`json:"id"` 
	Title			string	`json:"name"`
}

type QuestionAPI struct {
	Id 				bson.ObjectId 	`json:"_id" bson:"_id"`
	Body            string	`json:"body"`
	CreatedOn		time.Time  `json:"createdOn"`
}

//home page response type
type Home struct {
	SpaceAPIs		[]SpaceAPI `json:"followed_topics"`
	QuestionAPIs	[]QuestionAPI `json:"my_questions"`
}

// get space content from space server
type SpaceContentAPI struct {
	Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
    Title 		string 			`json:"title"`
    CreatedOn 	time.Time 		`json:"createdOn"`
    Description string 			`json:"description"`
    Tags 		[]SpaceTags 	`json:"tags"`
    //Questions	[]QuestionContentAPI  `json:"questions"`
}
// get Depth1 content from questions server
type Depth1 struct {
	Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
    Title 		string 			`json:"title"`
    CreatedOn 	time.Time 		`json:"createdOn"`
    Description string 			`json:"description"`
    Tags 		[]SpaceTags 	`json:"tags"`
    Questions	[]QuestionContentAPI  `json:"questions"`
}

type SpaceTags struct {
	Tag 		string			`json:"tag"`
}

// get question content from question server
type QuestionContentAPI struct {
    Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
	SpaceId 	bson.ObjectId 	`json:"spaceId" bson:"spaceId"`
    Body 		string 			`json:"body"`
    CreatedOn 	time.Time 		`json:"createdOn"`
    //CreatedBy 	bson.ObjectId 	`json:"createdBy" bson:"createdBy"`
    CreatedBy 	string	`json:"createdBy" bson:"createdBy"`
}

// insert data into mongo struct
type MongoInfoSpace struct {
	Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
	Title 		string 			`json:"title"`
	Questions 	[]SpaceQuestions 	`json:"questions"`
}

type SpaceQuestions struct {
	Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
}

var Spaces [] Space

