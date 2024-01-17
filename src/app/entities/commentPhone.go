package entities

import (
	"time"
)

type (
	// CommentsPhone
	CommentsPhone struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		PhoneID   uint      `json:"phone_id"`
		Name      string    `json:"name"`
		Comment   string    `json:"comment"`
		CreatedAt time.Time `json:"created_at"`
		UserID    uint      `json:"user_id"`
		User      User      `json:"-"`
		Phone     Phone     `json:"-"`
	}
)
