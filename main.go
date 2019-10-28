package main

import (
	"github.com/gorilla/mux"
	"log"
	"mockapi/handlers/contacts"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/contacts", contacts.Get).Methods("GET")
	router.HandleFunc("/contacts/{id}", contacts.Get).Methods("GET")
	router.HandleFunc("/contacts/{id}", contacts.Create).Methods("POST")
	router.HandleFunc("/contacts/{id}", contacts.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
