package main

import (
	"./handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	println("Welcome to the url shortener web service")

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.ShortenURL)
	router.HandleFunc("/{key}", handlers.RedirectURL)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
