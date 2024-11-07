package db

import (
	"log"
	"review-api/config"

	supabase "github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

// Connect initializes the Supabase client using SUPABASE_URL and SUPABASE_KEY
func Connect(cfg *config.Config) {
	client, err := supabase.NewClient(cfg.SupabaseURL, cfg.SupabaseKey, nil) // Passing `nil` for ClientOptions
	if err != nil {
		log.Fatalf("Failed to create Supabase client: %v", err)
	}
	SupabaseClient = client
	log.Println("Connected to Supabase database.")
}

// Disconnect sets SupabaseClient to nil, aligning with best practices
func Disconnect() {
	SupabaseClient = nil
	log.Println("Disconnected from Supabase database.")
}
