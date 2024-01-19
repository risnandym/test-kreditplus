package contract

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterOutput struct {
	UUID      uuid.UUID `gorm:"not null;unique" json:"uuid"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Phone     string    `gorm:"not null;" json:"phone"`
	LastLogin time.Time `gorm:"not null;" json:"last_login"`
}

func ValidateAndBuildUserRegister(c *gin.Context) (input RegisterInput, err error) {
	if err = c.ShouldBindJSON(&input); err != nil {
		return
	}
	return
}

func ValidateAndBuildUserLogin(c *gin.Context) (input LoginInput, err error) {
	if err = c.ShouldBindJSON(&input); err != nil {
		return
	}
	return
}
