package main

import (
	"Zametki/internal/requests"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/sign-up", requests.Registration).Methods("POST")
	router.HandleFunc("/sign-in", requests.Login).Methods("POST")
	router.HandleFunc("/notes", requests.GetNotes).Methods("GET")
	router.HandleFunc("/create-notes", requests.GetNotes).Methods("POST")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("error listen and serve: %v", err)
	}
}
