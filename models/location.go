package models

import "time"

type Location struct {
	CreatedAt time.Time `json:"created_at"`
	FullName  string    `json:"full_name"`
	ID        int       `json:"id"`
	ImageURL  string    `json:"image_url"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}
