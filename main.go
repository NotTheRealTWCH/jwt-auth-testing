package main

import (
	"jwt-auth-testing/db"
	"jwt-auth-testing/server"
	"log"
)

var host = "localhost"
var port = os.Getenv("PORT")

func main() {
	// init the DB
	db.InitDB()

	// start the server
	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting server!")
		log.Fatal(serverErr)
	}
}
