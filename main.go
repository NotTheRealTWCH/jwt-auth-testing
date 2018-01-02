package main

import (
	"jwt-auth-testing/db"
	"jwt-auth-testing/server"
	"log"
	"os"
)

var host = ""
var port = os.Getenv("PORT")

func main() {
	// init the DB
	db.InitDB()
	
	db.StoreUser("test", password, "user")

	// start the server
	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting server!")
		log.Fatal(serverErr)
	}
}
