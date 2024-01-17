package models

import (
	"time"
)

type (
	// News
	News struct {
		ID          uint           `json:"id" gorm:"primary_key"`
		Title       string         `json:"title"`
		Content     string         `json:"content"`
		Link_URL    string         `json:"link_url"`
		CreatedAt   time.Time      `json:"created_at"`
		UpdatedAt   time.Time      `json:"updated_at"`
		CreatorName string         `json:"creator_name"`
		UserID      uint           `json:"user_id"`
		User        User           `json:"-"`
		Comments    []CommentsNews `json:"-"`
	}
)
