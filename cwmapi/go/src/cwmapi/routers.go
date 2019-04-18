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
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Ping(w http.ResponseWriter, r *http.Request) {
	ping();
	fmt.Fprintf(w, "pong!")
}

var routes = Routes{
	Route{
		"ping",
		"GET",
		"/msgstore/v1/ping",
		Ping,
	},

	Route{
		"ObjectidsGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/objectids",
		ObjectidsGet,
	},

	Route{
		"ObjectidsGetSpaces",
		strings.ToUpper("Get"),
		"/msgstore/v1/objectids/spaces",
		ObjectidsGet,
	},

	Route{
		"ObjectidsGetQuestions",
		strings.ToUpper("Get"),
		"/msgstore/v1/objectids/questions/{spaceId}",
		ObjectidsGet,
	},

	Route{
		"ObjectidsGetAnswers",
		strings.ToUpper("Get"),
		"/msgstore/v1/objectids/answers/{questionId}",
		ObjectidsGet,
	},

	Route{
		"QuestionQuestionIdGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/question/{questionId}",
		QuestionQuestionIdGet,
	},

	Route{
		"QuestionsSpaceIdGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/questions/{spaceId}",
		QuestionsSpaceIdGet,
	},

	Route{
		"SpaceSpaceIdGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/space/{spaceId}",
		SpaceSpaceIdGet,
	},

	Route{
		"SpacesGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/spaces",
		SpacesGet,
	},

	{
		"AnswerAnswerIdGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/answer/{answerId}",
		AnswerAnswerIdGet,
	},

	{
		"AnswersQuestionIdGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/answers/{questionId}",
		AnswersQuestionIdGet,
	},

	{
		"CommentCommentIdGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/comment/{commentId}",
		CommentCommentIdGet,
	},

	{
		"CommentsAnswerIdGet",
		strings.ToUpper("Get"),
		"/msgstore/v1/comments/{answerId}",
		CommentsAnswerIdGet,
	},


}
