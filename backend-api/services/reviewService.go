package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"review-api/db"
	"review-api/models"
	"time"
)

func GetReviewsByCategoryAndDate(category, date string) ([]models.Review, error) {
	// Construct the Supabase REST API URL for the "reviews" table
	url := fmt.Sprintf("%s/rest/v1/reviews?category=eq.%s&review_date=eq.%s", os.Getenv("SUPABASE_URL"), category, date)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("apikey", os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code indicates an error
	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d\n", resp.StatusCode)
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Response body: %s\n", string(bodyBytes))
		return nil, fmt.Errorf("failed to fetch reviews: %s", resp.Status)
	}

	// Decode the JSON response into the reviews slice
	var reviews []models.Review
	if err := json.NewDecoder(resp.Body).Decode(&reviews); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Log and return the reviews
	log.Printf("Fetched %d reviews for category '%s' on date '%s'\n", len(reviews), category, date)
	return reviews, nil
}

type TrendData struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

// Get7DayTrend generates the count of reviews per day over the last 7 available days with data for a specific category
func Get7DayTrend(category, endDate string) ([]TrendData, error) {
	if db.SupabaseClient == nil {
		return nil, fmt.Errorf("database connection not initialized")
	}

	var trend []TrendData
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	// Calculate the start date for a 7-day trend
	startDate := end.AddDate(0, 0, -6)

	// Loop through each date in the range to get counts individually
	for i := 0; i < 7; i++ {
		date := startDate.AddDate(0, 0, i).Format("2006-01-02")

		var countResults []map[string]interface{}
		_, err := db.SupabaseClient.From("reviews").
			Select("id", "", true).
			Match(map[string]string{"category": category, "review_date": date}).
			ExecuteTo(&countResults)

		if err != nil {
			log.Printf("Error querying database for date %s: %v", date, err)
			continue
		}

		// Add the date and count of reviews for this day to the trend data
		trend = append(trend, TrendData{
			Date:  date,
			Count: len(countResults),
		})
	}

	// Log final trend data for debugging
	log.Printf("Final trend data for category '%s': %v", category, trend)

	return trend, nil
}
