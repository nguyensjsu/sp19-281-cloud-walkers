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

func getOr(filters []string, fieldType string)(bson.M){
	var query bson.M;

	switch len(filters) {

	case 0:

	case 1:
		query = bson.M{fieldType: bson.ObjectIdHex(filters[0])}
	default:
		orQuery := []bson.M{}

		for _, spaceId := range filters {
			if (bson.IsObjectIdHex(spaceId)) {
				orQuery = append(orQuery, bson.M{fieldType: bson.ObjectIdHex(spaceId)})
			}
		}

		query = bson.M{"$or": orQuery}
	}

	return query
}

func getAndFilters(leftSide bson.M, rightSide bson.M)(bson.M){
	if(leftSide == nil){
		return rightSide
	} else if (rightSide == nil){
		return leftSide
	}

	return bson.M{"$and" : []bson.M{leftSide, rightSide}}
}

func getSpaces (spaceFilter [] string, nestingLevel int) ([] Space){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_spaces)

	var query bson.M = getOr(spaceFilter, "_id");

	var spaceRecs []Space
	err = c.Find(query).All(&spaceRecs)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		for i := 0; i < len(spaceRecs); i++ {
			spaceRecs[i].Questions = getQuestions([]string{spaceRecs[i].Id.Hex()}, []string{}, nestingLevel - 1)
		}
	}

	return spaceRecs;
}

func getQuestions(spaceFilter [] string, questionFilter [] string, nestingLevel int) ([] Question){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	var query bson.M = getAndFilters(getOr(spaceFilter, "spaceId"), getOr(questionFilter, "_id"));
	var questionRecs []Question


	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_questions)

	fmt.Println(query)
	err = c.Find(query).All(&questionRecs)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		for i := 0; i < len(questionRecs); i++ {
			questionRecs[i].Answers = getAnswers([]string{questionRecs[i].Id.Hex()}, []string{}, nestingLevel - 1)
		}
	}

	return questionRecs;

}

func postQuestion(spaceId bson.ObjectId, newQ NewQuestion)(*Question){

	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_questions)

	qid := bson.NewObjectId()
	date := time.Now()

	var question = Question{
		Id: qid, SpaceId: spaceId, QuestionText: newQ.QuestionText, CreatedOn: date, CreatedBy: newQ.UserId}

	err = c.Insert(question)
	if err != nil {
		panic(err)
	}
	return &question
}

func getAnswers(questionFilter [] string, answerFilter [] string, nestingLevel int) ([] Answer){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}
	var query bson.M = getAndFilters(getOr(questionFilter, "questionId"), getOr(answerFilter, "_id"));
	var answerRecs []Answer

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_answers)

	fmt.Println(query)
	err = c.Find(query).All(&answerRecs)
	if err != nil {
		log.Fatal(err)
	}

	if(nestingLevel > 0){
		for i := 0; i < len(answerRecs); i++ {
			answerRecs[i].Comments = getComments([]string{answerRecs[i].Id.Hex()}, []string{}, nestingLevel - 1)
		}
	}

	return answerRecs;

}


func postAnswer(questionId bson.ObjectId, newA NewAnswer)(*Answer){

	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_answers)

	aid := bson.NewObjectId()
	date := time.Now()

	var answer = Answer{
		Id: aid, QuestionId: questionId, AnswerText: newA.AnswerText, CreatedOn: date, CreatedBy: newA.UserId}

	err = c.Insert(answer)
	if err != nil {
		panic(err)
	}
	return &answer
}



func getComments(answerFilter [] string, commentFilter [] string, nestingLevel int) ([] Comment){
	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}

	var answerQuery bson.M

	if(len(answerFilter) > 0){
		answerQuery = getAndFilters(getOr(answerFilter, "answerId"), bson.M{"parentCommentId": nil})
	}


	var query bson.M = getAndFilters(answerQuery, getOr(commentFilter, "_id"));
	var commentRecs []Comment

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

func postComment(answerId bson.ObjectId, newC NewComment)(*Comment){

	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_comments)

	aid := bson.NewObjectId()
	date := time.Now()

	var comment = Comment{
		Id: aid, AnswerId: answerId, CommentText: newC.CommentText, CreatedOn: date, CreatedBy: newC.UserId}

	err = c.Insert(comment)
	if err != nil {
		panic(err)
	}
	return &comment
}


func postReply(answerId bson.ObjectId, parentCommentId bson.ObjectId, newC NewComment)(*Comment){

	session, err := mgo.Dial(mongodb_server)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_comments)

	aid := bson.NewObjectId()
	date := time.Now()

	var comment = Comment{
		Id: aid, AnswerId: answerId, ParentCommentId: parentCommentId, CommentText: newC.CommentText, CreatedOn: date, CreatedBy: newC.UserId}

	err = c.Insert(comment)
	if err != nil {
		panic(err)
	}
	return &comment
}
