package models

import (
	"fmt"
	"time"
)

type Trip struct {
	Bus         Bus       `json:"bus"`
	BusID       int       `json:"bus_id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	From        Location  `json:"from"`
	FromID      int       `json:"from_id"`
	ID          int       `json:"id"`
	Price       string    `json:"price"`
	Quantity    int       `json:"quantity"`
	To          Location  `json:"to"`
	ToID        int       `json:"to_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (trip Trip) String() string {
	return fmt.Sprintf("%s to %s", trip.From.FullName, trip.To.FullName)
}
