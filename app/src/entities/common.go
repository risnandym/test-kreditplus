package entities

import "time"

type IDForm struct {
	ID uint `gorm:"primary_key" json:"id"`
}

type TimeStamp struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
