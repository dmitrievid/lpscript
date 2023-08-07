package main

import (
	"log"
	ws "logoper_services/postgres/script/webserv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error. Could not load environment variables: %v\n", err)
	}
	ws.HandleRequests()
}
