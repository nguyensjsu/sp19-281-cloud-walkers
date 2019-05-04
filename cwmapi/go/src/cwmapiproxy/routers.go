package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewServer() *negroni.Negroni {

	corsObj := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	n := negroni.Classic()
	mx := NewRouter()
	n.Use(corsObj)
	n.UseHandler(mx)
	return n
}

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


var routes = Routes{
	Route{
		"ping",
		"GET",
		"/cwmapiproxy/v1/ping",
		Ping,
	},

	Route{
		"ping",
		"GET",
		"/cwmapiproxy/v1/userFollow",
		UserFollows,
	},

	{
		"putUpdateCache",
		strings.ToUpper("Put"),
		"/msgstore/v1/flushcache",
		FlushCache,
	},

}
