package contract

import (
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ValidateAndBuildUserRegister(c *gin.Context) (input RegisterInput, err error) {
	if err = c.ShouldBindJSON(&input); err != nil {
		return
	}
	return
}
