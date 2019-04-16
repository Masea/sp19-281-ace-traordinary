package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUserEndpoint).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/users", handlers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/users", handlers.DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/users/{id}", handlers.UpdateUserEndpoint).Methods("PUT")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}