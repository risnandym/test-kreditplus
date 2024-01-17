package entities

import (
	"time"
)

type (
	// Brand
	Brand struct {
		ID          uint      `gorm:"primary_key" json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		UserID      uint      `json:"user_id"`
		User        User      `json:"-"`
		Phones      []Phone   `json:"-"`
	}
)
