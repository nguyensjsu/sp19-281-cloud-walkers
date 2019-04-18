/*
 * Cloud Walkers Message API
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

type Answer struct {
	Id  bson.ObjectId `json:"_id" bson:"_id"`
	QuestionId bson.ObjectId `json:"questionId" bson:"questionId"`
	Body string `json:"body"`
	CreatedOn time.Time `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
	Comments []Comment `json:"comments,omitempty"`
}
