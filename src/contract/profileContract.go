package contract

import (
	"test-kreditplus/src/entities"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileInput struct {
	AuthID       uint      `json:"auth_id,omitempty"`
	NIK          string    `json:"nik" validate:"required"`
	FullName     string    `json:"full_name" validate:"required"`
	LegalName    string    `json:"legal_name" validate:"required"`
	PlaceOfBirth string    `json:"place_of_birth" validate:"required"`
	DateOfBirth  time.Time `json:"date_of_birth" validate:"required"`
	Salary       float64   `json:"salary" validate:"required"`
	KtpImage     string    `json:"ktp_image" validate:"required"`
	SelfieImage  string    `json:"selfie_image" validate:"required"`
}

type ProfileOutput struct {
	ProfileInput
	Limit Limit `json:"limit"`
}

func ValidateAndBuildProfileInput(c *gin.Context) (request ProfileInput, err error) {
	authctx, _ := c.Get("auth")
	auth := authctx.(entities.Auth)

	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	request.AuthID = auth.ID
	return
}
