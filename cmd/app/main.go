package main

import (
	"Zametki/configs"
	"Zametki/database"
	"Zametki/internal/requests"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("error while init db: %v", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/sign-up", requests.Registration).Methods("POST")
	router.HandleFunc("/sign-in", requests.Login).Methods("POST")
	router.HandleFunc("/notes", requests.GetNotes).Methods("GET")
	router.HandleFunc("/create-notes", requests.GetNotes).Methods("POST")

	port := configs.GetEnv("SERVER_PORT")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("error listen and serve: %v", err)
	}
}
