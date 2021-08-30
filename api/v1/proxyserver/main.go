package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	jobsApiService := NewJobsApiService()
	jobsApiController := NewJobsApiController(jobsApiService)

	router := NewRouter(jobsApiController)

	log.Fatal(http.ListenAndServe(":4646", router))
}
