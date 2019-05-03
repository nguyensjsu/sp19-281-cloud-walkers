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
	"log"
	"os"
)

func main() {


	log.Printf("Server started")
	err := DbInit();

	if(err != nil){
		panic(err)
		return
	}
	ufInit()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := NewServer()
	server.Run(":" + port)

}
