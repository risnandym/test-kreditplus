package contract

import (
	"test-kreditplus/src/app/entities"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileInput struct {
	AuthID       int       `json:"auth_id"`
	NIK          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	PlaceOfBirth string    `json:"place_of_birth"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Salary       float32   `json:"salary"`
	KtpImage     string    `json:"ktp_image"`
	SelfieImage  string    `json:"selfie_image"`
}

func ValidateAndBuildProfileInput(c *gin.Context) (input ProfileInput, err error) {
	authctx, _ := c.Get("auth")
	auth := authctx.(entities.Auth)

	if err = c.ShouldBindJSON(&input); err != nil {
		return
	}

	input.AuthID = int(auth.ID)
	return
}
