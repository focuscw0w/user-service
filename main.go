package main

import (
	"net/http"
	"log"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request for '/'")
	})

	server := http.Server{
		Addr:   ":8080",
		Handler: router,
	}

	log.Println("Starting server on :8080")
	server.ListenAndServe()
}