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
	RandomAPIs		[]QuestionAPI `json:"random_questions"`
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

type HomeTest struct {
	TestTopic		[]TestTopic `json:"followed_topics"`
	QuestionAPIs	[]QuestionAPI `json:"my_questions"`
	RandomAPIs		[]QuestionAPI `json:"random_questions"`
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
    Body 		string 			`json:"questionText"`
    CreatedOn 	time.Time 		`json:"createdOn"`
    CreatedBy 	string	        `json:"createdBy" bson:"createdBy"`
    Topics      []TestTopic     `json:"topics"`
}

// get answer content from david
type AnswerContentAPI struct {
    Id  		bson.ObjectId 	`json:"_id" bson:"_id"`
    Body 		string 			`json:"answerText"`
    CreatedOn 	time.Time 		`json:"createdOn"`
    CreatedBy 	string	        `json:"createdBy" bson:"createdBy"`    
}

type TestTopic struct {
	Label	 string  `json:"label" bson:"label"`
}

/*** Deprecated
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

// db.u
type MUserSpace struct {
	UserId          	string   `json:"userId" bson:"userId"`	
	Uspaces		       []string  `json:"fspaces" bson:"fspaces"`
	Uquestions         []string  `json:"fquestions" bson:"fquestions"`
	UMquestions		   []string  `json:"myquestions" bson:"myquestions"`
	UAnswers		   []string	 `json:"manswers" bson:"manswers"`
}
*******/
/****************** Mongo Part**************************/
// db.uSpace
type MUserSpace struct {
	UserId          	string   `json:"userId" bson:"userId"`	
	Uspaces		        string  `json:"spaceId" bson:"spaceId"`
}
// db.uFQuestion
type MUserFQuestion struct {
	UserId          	string   `json:"userId" bson:"userId"`	
	FollowedQ           string  `json:"questionId" bson:"questionId"`
}
// db.uQuestion
type MUserQuestion struct {
	UserId          	string   `json:"userId" bson:"userId"`	
	Uquestions          string  `json:"questionId" bson:"questionId"`
}
// db.uAnswer
type MUserAnswer struct {
	UserId          	string   `json:"userId" bson:"userId"`	
	UAnswers		    string	 `json:"answerId" bson:"answerId"`
}

/****************** POST Part**************************/
type PostFollow struct {
	Action          	string   `json:"action" bson:"action"`	
	Id		        	string  `json:"id" bson:"id"`
	Unfollow			bool	`json:"unfollow" bson:"unfollow"`
}
// db.uFQuestion
type PostContent struct {
	Action         string   `json:"action" bson:"action"`	
	Id             string  `json:"id" bson:"id"`
}

type Success struct {
	Success		 bool	`json:"success" bson:"success"`
}

