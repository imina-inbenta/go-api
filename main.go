package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// fun main()
func main() {
	insertdata()
	router := mux.NewRouter()
	router.HandleFunc("/", DummyJSON).Methods("GET")
	port := os.Getenv("PORT") // Heroku provides the port to bind
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
