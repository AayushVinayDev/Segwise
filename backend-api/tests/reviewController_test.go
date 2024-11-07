package tests

import (
	"net/http"
	"net/http/httptest"
	"review-api/config"
	"review-api/db"
	"review-api/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReviewsEndpoint(t *testing.T) {
	// Set up database and router
	cfg := config.LoadConfig()
	db.Connect(cfg)
	defer db.Disconnect()
	router := routes.NewRouter()

	req, err := http.NewRequest("GET", "/reviews?category=Praises&date=2024-11-01", nil)
	assert.NoError(t, err)

	// Record the response
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Check the response
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Praises") // Check if the response body contains the expected category
}

func TestGetTrendEndpoint(t *testing.T) {
	// Set up database and router
	cfg := config.LoadConfig()
	db.Connect(cfg)
	defer db.Disconnect()
	router := routes.NewRouter()

	req, err := http.NewRequest("GET", "/trend?category=Praises&date=2024-11-01", nil)
	assert.NoError(t, err)

	// Record the response
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Check the response
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "date")  // Check if response contains a date field
	assert.Contains(t, recorder.Body.String(), "count") // Check if response contains a count field
}
