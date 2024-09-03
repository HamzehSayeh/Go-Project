package main

import (
	"github.com/HamzehSayeh/go-project/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Create a new router
	r := chi.NewRouter()

	// Define routes
	r.Post("/send-otp", handlers.SendOTP)

	// Start the server
	http.ListenAndServe(":8080", r)
}
