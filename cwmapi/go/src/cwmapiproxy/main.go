package main

import (
	"log"
	"os"
)

func main() {


	log.Printf("Server started")
	initProxy()
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := NewServer()
	server.Run(":" + port)

}

