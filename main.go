package main

import (
	"log"
	"net/http"
	"os"

	"golang-authentication-jwt/controllers"
	"golang-authentication-jwt/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// Protected routes
	s := router.PathPrefix("/api").Subrouter()
	s.Use(middlewares.JwtVerify)
	s.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the protected area!"))
	}).Methods("GET")

	log.Printf("Server started on :%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
