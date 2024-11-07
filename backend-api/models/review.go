package models

type Review struct {
	ID         string `json:"id"` // UUID as a string
	ReviewText string `json:"review_text"`
	ReviewDate string `json:"review_date"`
	Rating     int    `json:"rating"`
	Category   string `json:"category"`
}
