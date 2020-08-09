package models

import (
	"time"
)

// Category : using for inserting data or get data form database
type Category struct {
	ID        int8      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
