package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseURL string
	SupabaseKey string
}

func LoadConfig() *Config {
	// Attempt to load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using environment variables")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	// Check if essential variables are loaded
	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("Environment variables SUPABASE_URL or SUPABASE_KEY are missing")
	}

	return &Config{
		SupabaseURL: supabaseURL,
		SupabaseKey: supabaseKey,
	}
}
