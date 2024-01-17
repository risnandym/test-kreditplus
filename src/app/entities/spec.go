package entities

import (
	"time"
)

type (
	// Spec
	Spec struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		PhoneID   uint      `json:"phone_id"`
		Processor string    `json:"processor"`
		Memory    string    `json:"memory"`
		Storage   string    `json:"storage"`
		Screen    string    `json:"screen"`
		Camera    string    `json:"camera"`
		Price     string    `json:"price"`
		Review    string    `json:"review"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		UserID    uint      `json:"user_id"`
		User      User      `json:"-"`
		Phone     Phone     `json:"-"`
	}
)
