package entities

import (
	"time"
)

type (
	// CommentsNews
	CommentsNews struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		NewsID    uint      `json:"news_id"`
		Name      string    `json:"name"`
		Comment   string    `json:"comment"`
		CreatedAt time.Time `json:"created_at"`
		UserID    uint      `json:"user_id"`
		User      User      `json:"-"`
		News      News      `json:"-"`
	}
)
