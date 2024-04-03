package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-snowflake/routes"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("POST /query", routes.QueryHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
