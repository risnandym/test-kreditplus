package models

import (
	"time"
)

type (
	// Phone
	Phone struct {
		ID         uint            `json:"id" gorm:"primary_key"`
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
	}
)
