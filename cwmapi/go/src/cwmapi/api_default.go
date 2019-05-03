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
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// OK, this is obviously a no-no for real code.  Never put password in code
const jwtToken = "secret"

func getUserToken(w http.ResponseWriter, r *http.Request)(string, bool){
	tknStr := r.Header.Get("Authorization")

	if(len(tknStr) == 0){
		http.Error(w, "JWT User token required", http.StatusUnauthorized)
		return "", false

	}

	tokens := strings.Split(tknStr, " ")

	if(len(tokens) != 2){
		http.Error(w, "Expected type JWT not found", http.StatusUnauthorized)
		return "", false

	}

	return tokens[1], true
}

func getUserTokenFromRequest(w http.ResponseWriter, r *http.Request)(string, bool){
	// validate JWT Token

	tokenStr, ok := getUserToken(w, r)

	if(!ok){
		return "", false
	}

	tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtToken), nil
	})

	if tkn == nil || !tkn.Valid {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return "", false
	}

	claims, valid := tkn.Claims.(jwt.MapClaims)

	if err != nil || !valid{
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return "", false
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", false
	}


	//fmt.Println("claims", claims)

	return claims["id"].(string), true

}

func Ping(w http.ResponseWriter, r *http.Request) {
	/*_, ok := getUserTokenFromRequest(w, r)

	if(!ok){
		return
	}*/

	//fmt.Fprintf(w, "ping!")
	ping();
	jsonVal, _ := json.Marshal("pong");
	w.Write(jsonVal)
	w.WriteHeader(200)
	//fmt.Fprintf(w, "pong!\n")
}

func getUserFollowsTopics(w http.ResponseWriter, r *http.Request) ([]string, bool) {

	userId, ok := getUserTokenFromRequest(w, r)

	if(!ok){
		return []string{}, false
	}

	uToken, ok := getUserToken(w, r)
	if(!ok){
		return []string{}, false
	}

	followedTopics, err := getUserFollows(userId, uToken)
	if(err != nil){
		http.Error(w, err.Error(), 500)
		return []string{}, false
	}

	log.Println(followedTopics)
	return followedTopics, true

}

// GetTags - All the tags in the system
func GetTopics(w http.ResponseWriter, r *http.Request) {

	// is authorized user?

	_, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}

	pp := parseUrl(r.URL)
	filterMode := filterModeNormal

	_, needFavs := pp.queryParams["excludeFollowed"]
	var favorites []string

	if(needFavs ){
		// need to get favorites

		var ok bool
		favorites, ok = getUserFollowsTopics(w, r)

		if(!ok){
			return
		}

		filterMode = filterModeExclude
		log.Println("favorites: ", favorites)
	}

	topics := getTopics(favorites, filterMode,0);
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

// GetUserFeed - Gets questions for user based on followed topics
func GetUserFeed(w http.ResponseWriter, r *http.Request) {

	favorites, ok := getUserFollowsTopics(w, r)
	if(!ok){
		return;
	}
	var questions = []Question{}

	if(len(favorites) > 0) {
		questions = getQuestions([]string{}, favorites, intVal(parseUrl(r.URL), "depth", 0), nil, true)
	}

	jsonVal, err := json.MarshalIndent(questions, "", "   ");

	if(err != nil){
		log.Fatal(err);
	}
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonVal);
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {

	// is authorized user?
	_, ok := getUserTokenFromRequest(w, r)

	if(!ok){
		return;
	}

	pp := parseUrl(r.URL)

	first := pp.queryParams.Get("first")
	length := pp.queryParams.Get("length")

	topAnswer := booleanVal(pp, "topAnswer", false)

	var questions []Question

	if(len(first) > 0 || len(length) > 0) {
		if (len(first) == 0 || len(length) == 0) {
			http.Error(w, fmt.Sprintf("Invalid pagination.  Must use both skip and limit"), http.StatusBadRequest)
			return
		}

		pagination := Pagination{skip: ToInt(first), limit: ToInt(length)}
		questions = getQuestions(
			queryVals(pp.queryParams, "questionId"), queryVals(pp.queryParams, "topic"),
			intVal(pp, "depth", 0), &pagination, topAnswer)
	} else {
		questions = getQuestions(
			queryVals(pp.queryParams, "questionId"),
			queryVals(pp.queryParams, "topic"), intVal(pp, "depth", 0), nil, topAnswer)
	}
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

	uid, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}

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

	question := postQuestion(newQ, uid);

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

	// is authorized user?
	userId, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}

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

	questions := getQuestions([]string{questionId}, []string{}, 0, nil, false)

	if(questions == nil){
		http.Error(w, fmt.Sprintf("No question with answerId (%s) is found.", questionId), http.StatusNotFound)
		return
	}

	// only author can edit
	if(questions[0].CreatedBy != userId){
		http.Error(w, fmt.Sprintf("Invalid user.  Onlu author can edit."), http.StatusUnauthorized)
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

// DeleteQuestion - Delete one or more questions
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	_, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}

	pp := parseUrl(r.URL)

	deleteQuestions(queryVals(pp.queryParams, "questionId"))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetAnswers - Gets all answers for a question
func GetAnswers(w http.ResponseWriter, r *http.Request) {
	_, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}
	pp := parseUrl(r.URL)

	answers := getAnswers(
		queryVals(pp.queryParams, "questionId") ,
		queryVals(pp.queryParams, "answerId"), intVal(pp, "depth", 0))
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

	userId, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}


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

	question := getQuestions([]string{questionId}, []string{},0, nil, false)

	if(question == nil){
		http.Error(w, fmt.Sprintf("Invalid questionId: %s", questionId), http.StatusNotFound)
		return
	}


	answer := postAnswer(bson.ObjectIdHex(questionId), newA, userId);

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

	// is authorized user?
	userId, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}


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

	// only author can edit
	if(answers[0].CreatedBy != userId){
		http.Error(w, fmt.Sprintf("Invalid user.  Onlu author can edit."), http.StatusUnauthorized)
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

	comments := getComments(
		queryVals(pp.queryParams, "answerId"),
		queryVals(pp.queryParams, "commentId"), intVal(pp, "depth", 0))
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

	userId, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}



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


	comment := postComment(bson.ObjectIdHex(answerId), newC, userId);

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

	userId, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}



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

	comment := postReply(comments[0].AnswerId, bson.ObjectIdHex(commentId), newC, userId);

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

	// is authorized user?
	userId, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}


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

	// only author can edit
	if(comments[0].CreatedBy != userId){
		http.Error(w, fmt.Sprintf("Invalid user.  Onlu author can edit."), http.StatusUnauthorized)
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

// FlushCache - clear any caching for specific user.  Sent anytime a user makes a change, such as following/unfollowing a topic
func FlushCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// is authorized user?
	userId, success := getUserTokenFromRequest(w, r)

	if(!success){
		return;
	}

	flushCacheForUser(userId)

	w.WriteHeader(http.StatusNoContent)
}





