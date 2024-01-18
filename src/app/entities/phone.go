package entities

import (
	"time"
)

type Phone struct {
	IDForm
	Type       string          `json:"type"`
	Year       int             `json:"year"`
	BrandID    uint            `json:"brand_id"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	EditorName string          `json:"editor_name"`
	UserID     uint            `json:"user_id"`
	User       User            `json:"-"`
	Brand      Brand           `json:"-"`
	Comments   []CommentsPhone `json:"-"`
	TimeStamp
}
