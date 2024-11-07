package main

import (
	"log"
	"net/http"
	"review-api/config"
	"review-api/db"
	"review-api/routes"
)

func main() {
	// Load configuration from .env
	cfg := config.LoadConfig()

	// Initialize the Supabase client
	db.Connect(cfg)
	defer db.Disconnect()

	// Set up the router with all routes
	r := routes.NewRouter()

	// Serve static files for the UI from the "static" folder
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	// Start the HTTP server
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
