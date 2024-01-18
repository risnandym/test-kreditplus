package entities

import (
	"time"
)

type News struct {
	IDForm
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	Link_URL    string         `json:"link_url"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	CreatorName string         `json:"creator_name"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"-"`
	Comments    []CommentsNews `json:"-"`
	TimeStamp
}
