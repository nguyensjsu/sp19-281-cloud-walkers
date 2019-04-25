/*
 * Cloud Walkers Message API
 *
 * Post/read questions, answers, and comments
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
)


// GetTags - All the tags in the system
func GetTopics(w http.ResponseWriter, r *http.Request) {

	topics := getTopics([]string{}, 0);
	set := make(map[string]Topic)


	for k := range set {
		topics = append(topics, Topic{k})
	}

	jsonVal, err := json.MarshalIndent(topics, "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);
	w.Header().Set("Content-Type","application/json")

}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	pp := parseUrl(r.URL)

	questions := getQuestions(queryVals(pp.queryParams, "questionId"), queryVals(pp.queryParams, "topic"), depth(pp))
	jsonVal, err := json.MarshalIndent(questions, "", "   ");

	if(err != nil){
		log.Fatal(err);
	}
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);
}


// AddQuestion - post a question to a space
func PostQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var newQ NewQuestion

	err = json.Unmarshal(b, &newQ)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	question := postQuestion(newQ);

	jsonVal, err := json.MarshalIndent(question, "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);
}


// PutQuestionUpdate - update a question
func PutQuestionUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pp := parseUrl(r.URL)

	nIds := paramCount(pp.queryParams, "questionId")
	if(nIds != 1){
		http.Error(w, fmt.Sprintf("Only one questionId is allowed.  %d given", nIds), http.StatusBadRequest)
		return
	}

	questionId := pp.queryParams.Get("questionId")
	if(!bson.IsObjectIdHex(questionId)){
		http.Error(w, fmt.Sprintf("Invalid questionId:  %s", questionId), http.StatusBadRequest)
		return
	}

	questions := getQuestions([]string{questionId}, []string{}, 0)

	if(questions == nil){
		http.Error(w, fmt.Sprintf("No question with answerId (%s) is found.", questionId), http.StatusNotFound)
		return
	}

	var uo UpdateObject;

	err = json.Unmarshal(b, &uo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	questions[0].QuestionText = uo.Body

	putQuestionUpdate(questions[0])

	jsonVal, err := json.MarshalIndent(questions[0], "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);

}

// GetAnswers - Gets all answers for a question
func GetAnswers(w http.ResponseWriter, r *http.Request) {
	pp := parseUrl(r.URL)

	answers := getAnswers(queryVals(pp.queryParams, "questionId") ,queryVals(pp.queryParams, "answerId"), depth(pp))
	jsonVal, err := json.MarshalIndent(answers, "", "   ");

	if(err != nil){
		log.Fatal(err);
	}
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);
}

// PostAnswer - post an answer to a question
func PostAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pp := parseUrl(r.URL)

	nQuestionIds := paramCount(pp.queryParams, "questionId")
	if(nQuestionIds != 1){
		http.Error(w, fmt.Sprintf("Only one questionId is allowed.  %d given", nQuestionIds), http.StatusBadRequest)
		return
	}

	questionId := pp.queryParams.Get("questionId")
	if(!bson.IsObjectIdHex(questionId)){
		http.Error(w, fmt.Sprintf("Invalid questionId:  %s", questionId), http.StatusBadRequest)
		return
	}

	var newA NewAnswer

	err = json.Unmarshal(b, &newA)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	question := getQuestions([]string{questionId}, []string{},0)

	if(question == nil){
		http.Error(w, fmt.Sprintf("Invalid questionId: %s", questionId), http.StatusNotFound)
		return
	}


	answer := postAnswer(bson.ObjectIdHex(questionId), newA);

	jsonVal, err := json.MarshalIndent(answer, "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);

}

// PutAnswerUpdate - update an answer
func PutAnswerUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pp := parseUrl(r.URL)

	nIds := paramCount(pp.queryParams, "answerId")
	if(nIds != 1){
		http.Error(w, fmt.Sprintf("Only one answer is allowed.  %d given", nIds), http.StatusBadRequest)
		return
	}

	answerId := pp.queryParams.Get("answerId")
	if(!bson.IsObjectIdHex(answerId)){
		http.Error(w, fmt.Sprintf("Invalid answerId:  %s", answerId), http.StatusBadRequest)
		return
	}

	answers := getAnswers([]string{}, []string{answerId}, 0)

	if(answers == nil){
		http.Error(w, fmt.Sprintf("No answer with answerId (%s) is found.", answerId), http.StatusNotFound)
		return
	}

	var uo UpdateObject;

	err = json.Unmarshal(b, &uo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	answers[0].AnswerText = uo.Body

	putAnswerUpdate(answers[0])

	jsonVal, err := json.MarshalIndent(answers[0], "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);

}


// GetComments - Gets all top-level comments for an answer
func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pp := parseUrl(r.URL)

	comments := getComments(queryVals(pp.queryParams, "answerId"), queryVals(pp.queryParams, "commentId"), depth(pp))
	jsonVal, err := json.MarshalIndent(comments, "", "   ");

	if(err != nil){
		log.Fatal(err);
	}
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);
}


// PostComment - post a comment to an answer
func PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pp := parseUrl(r.URL)

	nAnswerIds := paramCount(pp.queryParams, "answerId")
	if(nAnswerIds != 1){
		http.Error(w, fmt.Sprintf("Only one answer is allowed.  %d given", nAnswerIds), http.StatusBadRequest)
		return
	}

	answerId := pp.queryParams.Get("answerId")
	if(!bson.IsObjectIdHex(answerId)){
		http.Error(w, fmt.Sprintf("Invalid answerId:  %s", answerId), http.StatusBadRequest)
		return
	}

	var newC NewComment

	err = json.Unmarshal(b, &newC)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	answer := getAnswers([]string{}, []string{answerId}, 0)

	if(answer == nil){
		http.Error(w, fmt.Sprintf("Invalid answerId: %s", answerId), http.StatusNotFound)
		return
	}


	comment := postComment(bson.ObjectIdHex(answerId), newC);

	jsonVal, err := json.MarshalIndent(comment, "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);

}

// PostComment - post a comment to an answer
func PostReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pp := parseUrl(r.URL)

	nCommentIds := paramCount(pp.queryParams, "commentId")
	if(nCommentIds != 1){
		http.Error(w, fmt.Sprintf("Only one comment is allowed.  %d given", nCommentIds), http.StatusBadRequest)
		return
	}

	commentId := pp.queryParams.Get("commentId")
	if(!bson.IsObjectIdHex(commentId)){
		http.Error(w, fmt.Sprintf("Invalid commentId:  %s", commentId), http.StatusBadRequest)
		return
	}

	var newC NewComment

	err = json.Unmarshal(b, &newC)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	comments := getComments([]string{}, []string{commentId}, 0)

	if comments == nil{
		http.Error(w, fmt.Sprintf("No comment with commentId (%s) is found.", commentId), http.StatusNotFound)
		return
	}

	comment := postReply(comments[0].AnswerId, bson.ObjectIdHex(commentId), newC);

	jsonVal, err := json.MarshalIndent(comment, "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);

}

// PutCommentUpdate - update a comment or reply
func PutCommentUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pp := parseUrl(r.URL)

	nCommentIds := paramCount(pp.queryParams, "commentId")
	if(nCommentIds != 1){
		http.Error(w, fmt.Sprintf("Only one comment is allowed.  %d given", nCommentIds), http.StatusBadRequest)
		return
	}

	commentId := pp.queryParams.Get("commentId")
	if(!bson.IsObjectIdHex(commentId)){
		http.Error(w, fmt.Sprintf("Invalid commentId:  %s", commentId), http.StatusBadRequest)
		return
	}

	comments := getComments([]string{}, []string{commentId}, 0)

	if(comments == nil){
		http.Error(w, fmt.Sprintf("No comment with commentId (%s) is found.", commentId), http.StatusNotFound)
		return
	}

	var uo UpdateObject;

	err = json.Unmarshal(b, &uo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	comments[0].CommentText = uo.Body

	putComentUpdate(comments[0])

	jsonVal, err := json.MarshalIndent(comments[0], "", "   ");

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);
}





