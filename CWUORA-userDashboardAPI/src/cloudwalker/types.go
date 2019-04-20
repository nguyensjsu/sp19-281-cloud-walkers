package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)



// type Space struct {
// 	SpaceId         string     `json:"spaceId"` 
// 	Title			string	`json:"title"`
// 	CreatedOn      time.Time `json:"createdOn"`
// 	Description		string  `json:"description"`
// }

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

/******************Frontend Part**************************/
//home page response type
type Home struct {
	SpaceAPIs		[]SpaceAPI `json:"followed_topics"`
	QuestionAPIs	[]QuestionAPI `json:"my_questions"`
}

type SpaceAPI struct {
	Id 				bson.ObjectId 	`json:"_id" bson:"_id"`
	Title			string	`json:"name"`
}

type QuestionAPI struct {
	Id 				bson.ObjectId 	`json:"_id" bson:"_id"`
	Body            string	`json:"body"`
	CreatedOn		time.Time  `json:"createdOn"`
}

// get space content from space server
// type SpaceContentAPI struct {
// 	Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
//     Title 		string 			`json:"title"`
//     CreatedOn 	time.Time 		`json:"createdOn"`
//     Description string 			`json:"description"`
//     Tags 		[]SpaceTags 	`json:"tags"`
// }

/******************David Part**************************/
type SpaceContentAPI struct {
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

// get question content from david
type QuestionContentAPI struct {
    Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
	SpaceId 	bson.ObjectId 	`json:"spaceId" bson:"spaceId"`
    Body 		string 			`json:"questionText"`
    CreatedOn 	time.Time 		`json:"createdOn"`
    CreatedBy 	string	`json:"createdBy" bson:"createdBy"`
    Answers	    []AnswerContentAPI  `json:"answers"`
}

// get answer content from david
type AnswerContentAPI struct {
    Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
	QuestionId 	bson.ObjectId 	`json:"questionId" bson:"questionId"`
    Body 		string 			`json:"answerText"`
    CreatedOn 	time.Time 		`json:"createdOn"`
    CreatedBy 	string	`json:"createdBy" bson:"createdBy"`
}
/****************** Mongo Part**************************/

// db.space
type MSpace struct {
	Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
	Title 		string 			`json:"title"`
	Questions 	[]MQuestion 	`json:"questions"`
}
// db.question
type MQuestion struct {
	Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
	Body 		string			`json:"questionText"`
	CreatedOn 	time.Time 		`json:"createdOn"`
	Answers     []MAnswer		`json:"answers"`
}
// db.answer
type MAnswer struct {
	Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
	Body 		string			`json:"answerText"`
	CreatedOn 	time.Time 		`json:"createdOn"`
}
// db.user
type MUserProfile struct {
	userId          	string    `json:"userId" bson:"userId"`	
	uspaces		        []SpaceAPI  `json:"followedspaces" bson:"followedspaces"`
	uquestions          []QuestionAPI `json:"followedquestions" bson:"followedquestions"`
}

