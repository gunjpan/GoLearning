package main

import (
	"log"
	"net/http"

	"github.com/GoLearning/go-server/models"
)

func main() {
	log.Printf("Server started at localhost:8080")
	models.CreateDummyLots()
	router := models.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
