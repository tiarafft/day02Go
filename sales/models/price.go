package models

import "time"

// Price : table on db
type Price struct {
	ID        int8      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
