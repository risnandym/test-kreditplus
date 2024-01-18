package entities

import (
	"time"
)

type CommentsNews struct {
	IDForm
	NewsID    uint      `json:"news_id"`
	Email     string    `json:"email"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"-"`
	News      News      `json:"-"`
	TimeStamp
}
