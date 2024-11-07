package tests

import (
	"log"
	"review-api/config"
	"review-api/db"
	"review-api/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetReviewsByCategoryAndDate(t *testing.T) {
	// Load config and initialize DB for testing
	cfg := config.LoadConfig()
	db.Connect(cfg)
	defer db.Disconnect()

	// Test case: valid category and date with existing reviews
	category := "Praises"
	date := time.Now().Format("2006-01-02") // use today's date for test
	reviews, err := services.GetReviewsByCategoryAndDate(category, date)
	assert.NoError(t, err)
	if len(reviews) > 0 {
		assert.Equal(t, category, reviews[0].Category)
	}

	// Test case: category and date with no reviews
	nonexistentCategory := "Nonexistent"
	reviews, err = services.GetReviewsByCategoryAndDate(nonexistentCategory, date)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(reviews))
}

func TestGet7DayTrend(t *testing.T) {
	// Load config and initialize DB for testing
	cfg := config.LoadConfig()
	db.Connect(cfg)
	defer db.Disconnect()

	// Test case: valid category with some trend data
	category := "Praises"
	endDate := time.Now().Format("2006-01-02") // use today's date for test
	trendData, err := services.Get7DayTrend(category, endDate)
	assert.NoError(t, err)

	if len(trendData) > 0 {
		for _, trend := range trendData {
			assert.NotEmpty(t, trend.Date)
			assert.GreaterOrEqual(t, trend.Count, 0)
		}
	} else {
		log.Println("No trend data available for the test category; adjust the category or date as needed.")
	}
}
