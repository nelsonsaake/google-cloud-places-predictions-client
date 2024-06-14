package models

import "time"

type Bus struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	ImageURL  string    `json:"image_url"`
	Name      string    `json:"name"`
	NameAbbr  string    `json:"name_abbr"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}
