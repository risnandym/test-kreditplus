package entities

import (
	"time"
)

type Profile struct {
	IDForm
	AuthID       uint      `json:"auth_id"`
	NIK          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	PlaceOfBirth string    `json:"place_of_birth"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Salary       float64   `json:"salary"`
	KtpImage     string    `json:"ktp_image"`
	SelfieImage  string    `json:"selfie_image"`
	TimeStamp
}
