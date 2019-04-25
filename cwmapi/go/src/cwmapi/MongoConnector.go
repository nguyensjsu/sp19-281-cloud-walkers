package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"strings"
	"time"
)


var mongodb_server = "localhost"
var mongodb_database = "cmpe281"
var collection_spaces = "spaces"
var collection_questions = "questions"
var collection_answers = "answers"
var collection_comments = "comments"
var collection_topics = "topics"
var db_user string //= "admin"
var db_pwd string //= "query"

var server1 string // = "52.10.186.169" // set in environment
var server2 string // = "54.148.30.107" // set in environment
var server3 string // = "54.203.107.100" // set in environment

var dialInfo = &mgo.DialInfo{}

var topicChan chan Topic

type Pagination struct{
	skip int
	limit int
}

func DbInit(){
	server1 = os.Getenv("MONGO1")
	server2 = os.Getenv("MONGO2")
	server3 = os.Getenv("MONGO3")
	db_user = os.Getenv("MONGO_ADMIN")
	db_pwd = os.Getenv("MONGO_PASSWORD")

	if(len(server1) == 0 && len(server2) == 0 && len(server3) == 0){
		server1 = mongodb_server
		log.Println("Connected to local host")
	} else {
		log.Printf("M1: {%s} M2: {%s} M3: {%s} => Admin: {%s} PWD: {%s}", server1, server2, server3, db_user, db_pwd)
	}

	dialInfo = &mgo.DialInfo{
		Addrs:    []string{server1, server2, server3},
		Timeout:  30 * time.Second,
		Database: "admin",
		Username: db_user,
		Password: db_pwd,
	}

	topicChan = make(chan Topic, 1000)

	go syncTopics(topicChan)
}

/**
	thread to keep topics in sync with question topics (add, but never delete)
 */
func syncTopics(topicChan <-chan Topic){

	for topic := range topicChan {
		if(len(getTopics([]string{topic.Label}, 0)) == 0){
			updateTopic(topic)
		}
	}
}


func dial() (*mgo.Session, error){
	return mgo.DialWithInfo(dialInfo)
	//return mgo.Dial(mongodb_server)
}

func getOr_id(filters []string, fieldType string)(bson.M){
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

func getAnd_str(filters []string, fieldType string)(bson.M){
	var query bson.M;

	switch len(filters) {

	case 0:

	case 1:
		query = bson.M{fieldType: filters[0]}
	default:
		andQuery := []bson.M{}

		for _, spaceId := range filters {
			andQuery = append(andQuery, bson.M{fieldType: spaceId})
		}

		query = bson.M{"$and": andQuery}
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

func getOrFilters(leftSide bson.M, rightSide bson.M)(bson.M){
	if(leftSide == nil){
		return rightSide
	} else if (rightSide == nil){
		return leftSide
	}

	return bson.M{"or" : []bson.M{leftSide, rightSide}}
}


func ping(){
	_, err := dial()
	if err != nil {
		panic(err)
	}
}

func getQuestions(questionFilter [] string, topicFilter []string, nestingLevel int, paginate *Pagination) ([] Question){
	session, err := dial()
	if err != nil {
		panic(err)
	}

	questionQuery := getOr_id(questionFilter, "_id");
	var topicAndQueries []bson.M

	for _, topic := range topicFilter {
		var andQuery []bson.M
		for _, andTopic := range strings.Split(topic, ","){
			andQuery = append(andQuery, bson.M{"topics": bson.M{"$elemMatch": bson.M{"label": andTopic}}})
		}
		topicAndQueries = append(topicAndQueries, bson.M{"$and": andQuery})
	}

	var topicOrQuery bson.M

	if(topicOrQuery != nil){
		topicOrQuery = bson.M{"$or": topicAndQueries}
	}

	query := getOrFilters(questionQuery, topicOrQuery)


	var questionRecs []Question


	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(collection_questions)

	fmt.Println(query)

	if(paginate != nil){
		fmt.Println("Paginate: ", paginate)
		err = c.Find(query).Skip(paginate.skip).Limit(paginate.limit).All(&questionRecs)
	} else {
		err = c.Find(query).All(&questionRecs)
	}

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

func postQuestion(newQ NewQuestion)(*Question){

	session, err := dial()
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_questions)

	qid := bson.NewObjectId()
	date := time.Now()

	var question = Question{
		Id: qid, QuestionText: newQ.QuestionText, CreatedOn: date, CreatedBy: newQ.UserId, Topics: newQ.Topics}

	err = c.Insert(question)
	if err != nil {
		panic(err)
	}

	for _, topic := range question.Topics{
		topicChan <- topic
	}
	return &question
}


func putQuestionUpdate(question Question){
	session, err := dial()
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_questions)

	err = c.Update(bson.M{"_id" : question.Id}, question )
	if err != nil {
		panic(err)
	}
}

func getAnswers(questionFilter [] string, answerFilter [] string, nestingLevel int) ([] Answer){
	session, err := dial()
	if err != nil {
		panic(err)
	}
	var query bson.M = getAndFilters(getOr_id(questionFilter, "questionId"), getOr_id(answerFilter, "_id"));
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

	session, err := dial()
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

func putAnswerUpdate(answer Answer){
	session, err := dial()
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_answers)

	err = c.Update(bson.M{"_id" : answer.Id}, answer )
	if err != nil {
		panic(err)
	}
}



func getComments(answerFilter [] string, commentFilter [] string, nestingLevel int) ([] Comment){
	session, err := dial()
	if err != nil {
		panic(err)
	}

	var answerQuery bson.M

	if(len(answerFilter) > 0){
		answerQuery = getAndFilters(getOr_id(answerFilter, "answerId"), bson.M{"parentCommentId": nil})
	}


	var query bson.M = getAndFilters(answerQuery, getOr_id(commentFilter, "_id"));
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
	session, err := dial()
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

	session, err := dial()
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

func putComentUpdate(comment Comment){
	session, err := dial()
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_comments)


	err = c.Update(bson.M{"_id" : comment.Id}, comment )
	if err != nil {
		panic(err)
	}
}

func postReply(answerId bson.ObjectId, parentCommentId bson.ObjectId, newC NewComment)(*Comment){

	session, err := dial()
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

func getTopics(labelFilters []string, nestingLevel int ) []Topic {
	session, err := dial()
	if err != nil {
		panic(err)
	}

	var topicRecs []Topic

	var query bson.M

	if(len(labelFilters) == 1){
		query = bson.M{"label": labelFilters[0]}

	} else{
		orQuery := []bson.M{}

		for _, label := range labelFilters {
				orQuery = append(orQuery, bson.M{"label": label})

		}
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_topics)


	err = c.Find(query).All(&topicRecs)
	if err != nil {
		log.Fatal(err)
	}

	return topicRecs;
}

func updateTopic(topic Topic)(Topic){
	session, err := dial()
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(collection_topics)

	chg, err := c.Upsert(bson.M{"label" : topic.Label}, topic)

	log.Printf("%s/n", chg)

	return topic
}

