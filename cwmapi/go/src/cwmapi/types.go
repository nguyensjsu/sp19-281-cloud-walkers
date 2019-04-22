/*
 * CWMAPI
 *
 * Post/read questions, answers, and comments
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Topic struct {

	Label string `json:"label" bson:"label"`
}

type Question struct {

	Id  bson.ObjectId `json:"_id" bson:"_id"`

	QuestionText string `json:"questionText" bson:"questionText"`

	CreatedOn time.Time `json:"createdOn" bson:"createdOn"`

	CreatedBy string `json:"createdBy"`

	Topics []Topic `json:"topics,omitempty" bson:"topics"`

	Answers []Answer `json:"answers,omitempty"`
}

type NewQuestion struct {

	UserId string `json:"userId"`

	QuestionText string `json:"questionText" bson:"questionText"`

	Topics []Topic `json:"topics,omitempty"`}

type Answer struct {
	Id  bson.ObjectId `json:"_id" bson:"_id"`
	QuestionId bson.ObjectId `json:"questionId" bson:"questionId"`
	AnswerText string `json:"answerText" bson:"answerText"`
	CreatedOn time.Time `json:"createdOn" bson:"createdOn"`
	CreatedBy string `json:"createdBy"  bson:"createdBy"`
	Comments []Comment `json:"comments,omitempty"`
}

type NewAnswer struct {

	UserId string `json:"userId"`

	AnswerText string `json:"answerText" bson:"answerText"`
}

type Comment struct {

	Id  bson.ObjectId `json:"_id" bson:"_id"`

	AnswerId bson.ObjectId `json:"answerId" bson:"answerId"`

	// The parent comment is for comment trees.  For first-gen comment, the parent is answer, and there is no parent comment.
	ParentCommentId bson.ObjectId `json:"parentCommentId,omitempty" bson:"parentCommentId,omitempty"`

	CommentText string `json:"commentText" bson:"commentText"`

	CreatedOn time.Time `json:"createdOn" bson:"createdOn"`

	CreatedBy string `json:"createdBy"  bson:"createdBy"`

	Replies []Comment `json:"replies,omitempty"`
}

type NewComment struct {

	UserId string `json:"userId"`

	CommentText string `json:"commentText,omitempty" bson:"commentText"`
}

type UpdateObject struct {

	Body string `json:"body,omitempty"`
}