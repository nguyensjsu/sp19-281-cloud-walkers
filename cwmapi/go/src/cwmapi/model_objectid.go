package main

import (
	"gopkg.in/mgo.v2/bson"
)

type ObjectId struct {

	Id  bson.ObjectId `json:"_id" bson:"_id"`
}
