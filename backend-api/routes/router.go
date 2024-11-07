package routes

import (
	"review-api/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/reviews", controllers.GetReviews).Methods("GET")
	r.HandleFunc("/trend", controllers.GetTrend).Methods("GET")

	return r
}
