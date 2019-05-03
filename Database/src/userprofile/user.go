package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	User struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Email     string        `json:"email" bson:"email"`
		Password  string        `json:"password,omitempty" bson:"password"`
		FirstName	string		`json:"firstname" bson:"firstname"`
		LastName	string		`json:"lastname" bson:"lastname"`
		Token     string        `json:"token,omitempty" bson:"-"`
	}

	SignUpResponse struct {
		FirstName	string		`json:"firstname" bson:"firstname"`
		LastName	string		`json:"lastname" bson:"lastname"`
		Message 	string		`json:"Message" bson:"-"`
	}

	LoginResponse struct {
		FirstName	string		`json:"firstname" bson:"firstname"`
		LastName	string		`json:"lastname" bson:"lastname"`
		Token     string        `json:"token,omitempty" bson:"-"`
	}

	UserRecord struct {
		FirstName	string		`json:"firstname" bson:"firstname"`
		LastName	string		`json:"lastname" bson:"lastname"`
		LoginTime	time.Time	`json:"Login Time" bson:"Login Time"`
	}

	UserRecordAPI struct {
		AllRecords	[]UserRecord `json:"User Records"`
	}
)




























































