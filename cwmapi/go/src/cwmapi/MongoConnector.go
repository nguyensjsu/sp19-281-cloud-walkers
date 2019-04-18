package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"

)

var mongodb_server = "localhost"
var mongodb_database = "cmpe281"
var collection_spaces = "spaces"
var collection_questions = "questions"
var collection_answers = "answers"
var collection_comments = "comments"
var db_user = "admin"
var db_pwd = "query"

var server1 = "52.10.186.169" // set in environment
var server2 = "54.148.30.107" // set in environment
var server3 = "54.203.107.100" // set in environment

func DbInit(local bool){

	if(!local){
		//server1 = os.Getenv("MONGO1")
		//server2 = os.Getenv("MONGO2")
		//server3 = os.Getenv("MONGO3")
		//mongodb_server = server1
	}
}

func dialInfo()(*mgo.DialInfo){
	return &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  30 * time.Second,
		Database: mongodb_database,
		Username: db_user,
		Password: db_pwd,
	}
}



func ping(){
	_, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
}

func getObjectIds(ids []string) ([]bson.ObjectId){

	var ret []bson.ObjectId;

	switch ids[0] {
	case "":
		for i := 0; i < 10; i++ {
			ret = append(ret, bson.NewObjectId())
		}
	case "spaces":
		spaces := getSpaces(0);
		for i := 1; i < len(spaces); i++{
			ret = append(ret, spaces[i].Id)
		}
	case "questions":
		questions := getQuestions(ids[1], 0)
		for i := 0; i < len(questions); i++ {
			ret = append(ret, questions[i].Id)
		}
	case "answers":
		answers := getAnswers(ids[1], 0)
		for i := 0; i < len(answers); i++ {
			ret = append(ret, answers[i].Id)
		}
	}

	return ret;
}

func getSpaces (nestingLevel int) ([] Space){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_spaces)
	var spaceRecs []Space
	err = c.Find(nil).All(&spaceRecs)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		for i := 0; i < len(spaceRecs); i++ {
			spaceRecs[i].Questions = getQuestions(spaceRecs[i].Id.Hex(), nestingLevel - 1)
		}
	}

	return spaceRecs;
}

func getSpace (spaceId string, nestingLevel int) (Space){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	query := bson.M{"_id": bson.ObjectIdHex(spaceId)}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_spaces)
	fmt.Println(query)
	var spaceRec Space
	err = c.Find(query).One(&spaceRec)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		spaceRec.Questions = getQuestions(spaceId, nestingLevel -1)
	}

	return spaceRec;
}

func getQuestions(spaceId string, nestingLevel int) ([] Question){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	var query bson.M;
	var questionRecs []Question

	if(bson.IsObjectIdHex(spaceId)) {
		query = bson.M{"spaceId": bson.ObjectIdHex(spaceId)}
	} else {
		return questionRecs
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_questions)

	fmt.Println(query)
	err = c.Find(query).All(&questionRecs)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		for i := 0; i < len(questionRecs); i++ {
			questionRecs[i].Answers = getAnswers(questionRecs[i].Id.Hex(), nestingLevel - 1)
		}
	}

	return questionRecs;

}

func getQuestion(questionId string, nestingLevel int) (Question){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	query := bson.M{"_id": bson.ObjectIdHex(questionId)}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_questions)
	var questionRec Question
	fmt.Println(query)
	err = c.Find(query).One(&questionRec)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		questionRec.Answers = getAnswers(questionId, nestingLevel - 1)
	}

	return questionRec;
}

func getAnswers(questionId string, nestingLevel int) ([] Answer){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	var query bson.M;
	var answerRecs []Answer

	// should return error if not valid id
	if(bson.IsObjectIdHex(questionId)) {
		query = bson.M{"questionId": bson.ObjectIdHex(questionId)}
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_answers)

	fmt.Println(query)
	err = c.Find(query).All(&answerRecs)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		for i := 0; i < len(answerRecs); i++ {
			answerRecs[i].Comments = getCommentsForAnswer(answerRecs[i].Id.Hex(), nestingLevel - 1)
		}
	}

	return answerRecs;

}

func getAnswer(answerId string, nestingLevel int) (Answer){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	var query bson.M
	var answerRec Answer

	// should return error if not valid id
	if(bson.IsObjectIdHex(answerId)) {
		query = bson.M{"_id": bson.ObjectIdHex(answerId)}
	} else{
		return answerRec
	}


	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_answers)

	fmt.Println(query)
	err = c.Find(query).One(&answerRec)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		answerRec.Comments = getCommentsForAnswer(answerId, nestingLevel - 1)
	}
	return answerRec;
}


func getCommentsForAnswer(answerId string, nestingLevel int) ([] Comment){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	var query bson.M;
	var commentRecs []Comment

	// should return error if not valid id
	if(bson.IsObjectIdHex(answerId)) {

		query = bson.M{
						"$and": []bson.M{ // you can try this in []interface
						bson.M{"answerId": bson.ObjectIdHex(answerId)},
						bson.M{"parentCommentId": nil}},
						}
	}else {
		return commentRecs
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_comments)

	fmt.Println(query)
	err = c.Find(query).All(&commentRecs)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		for i := 0; i < len(commentRecs); i++{
			commentRecs[i].Replies = getCommentChildren(commentRecs[i].Id, nestingLevel - 1)
		}
	}
	return commentRecs;

}

func getCommentChildren(commentId bson.ObjectId, nestingLevel int ) ([]Comment){
	var commentRecs []Comment
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	var query = bson.M{"parentCommentId": commentId};

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_comments)

	fmt.Println(query)
	err = c.Find(query).All(&commentRecs)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0) {
		for i := 0; i < len(commentRecs); i++ {
			children := getCommentChildren(commentRecs[i].Id, nestingLevel - 1)
			for j := 0; j < len(children); j++ {
				commentRecs[i].Replies = append(commentRecs[i].Replies, children[j]);
			}
		}
	}
	return commentRecs
}

func getComment(commentId string, nestingLevel int) (Comment){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	var query bson.M
	var commentRec Comment

	// should return error if not valid id
	if(bson.IsObjectIdHex(commentId)) {
		query = bson.M{"_id": bson.ObjectIdHex(commentId)}
	} else{
		return  commentRec
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_comments)

	fmt.Println(query)
	err = c.Find(query).One(&commentRec)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		commentRec.Replies = getCommentChildren(commentRec.Id, nestingLevel - 1)

	}

	return commentRec;
}





