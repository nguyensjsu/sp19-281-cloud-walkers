package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

//type (
//	Handler struct {
//		DB *mgo.Session
//	}
//)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func main() {
	// echo instance
	e := echo.New()

	e.Logger.SetLevel(log.ERROR)

	// returns a middleware that logs HTTP requests
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet,http.MethodPut,http.MethodPost,http.MethodDelete},
		AllowHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}))

	// return a JWT auth middleware with config
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{

		SigningKey: []byte(Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	// Database connection
	db, err := mgo.Dial("54.201.232.46:27017")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Create indices
	if err = db.Copy().DB("cw_user").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"_id"},
		//Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	//Initialize handler
	h := &Handler{DB: db}

	//var h *Handler
	//h.DB = db

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}

