package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"review-api/services"
)

// GetReviews handles GET requests for fetching reviews by category and date
func GetReviews(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	category := r.URL.Query().Get("category")
	date := r.URL.Query().Get("date")

	// Log received parameters for debugging
	log.Printf("Received request - Category: %s, Date: %s\n", category, date)

	// Fetch data from service layer
	reviews, err := services.GetReviewsByCategoryAndDate(category, date)
	if err != nil {
		log.Printf("Error fetching reviews: %v\n", err)
		http.Error(w, "Error fetching reviews", http.StatusInternalServerError)
		return
	}

	// Log the count of reviews fetched
	log.Printf("Fetched %d reviews\n", len(reviews))

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the reviews into JSON and write to the response
	if err := json.NewEncoder(w).Encode(reviews); err != nil {
		log.Printf("Error encoding response to JSON: %v\n", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// GetTrend handles GET requests for fetching 7-day trend data by category
func GetTrend(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	endDate := r.URL.Query().Get("date")

	// Check if the endDate parameter is provided
	if endDate == "" {
		http.Error(w, "Missing or invalid date parameter", http.StatusBadRequest)
		return
	}

	log.Printf("Received request for 7-day trend - Category: %s, End Date: %s\n", category, endDate)

	// Fetch trend data from the service layer
	trend, err := services.Get7DayTrend(category, endDate)
	if err != nil {
		log.Printf("Error fetching trend data: %v\n", err)
		http.Error(w, "Error fetching trend data", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched 7-day trend data for category '%s'\n", category)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(trend); err != nil {
		log.Printf("Error encoding response to JSON: %v\n", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
