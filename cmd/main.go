package main

import (
	"fmt"
	"log"
	routes "recruitment-test/application/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("server start on port 9000")
	routes.RunServer()
}
