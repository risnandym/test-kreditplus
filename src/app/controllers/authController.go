package controllers

import (
	"errors"
	"final-project/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChangePasswordInput struct {
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// LoginUser godoc
// @Summary Login.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	user := map[string]string{
		"username": u.Username,
		"email":    u.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "username": user, "token": token})

}

// Register godoc
// @Summary Register a User.
// @Description registering a user from public access.
// @Tags Auth
// @Param Body body RegisterInput true "the body to register a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password
	u.FullAccess = false

	_, err := u.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var FA string

	if u.FullAccess == false {
		FA = "false"
	}

	user := map[string]string{

		"username":    input.Username,
		"email":       input.Email,
		"full_access": FA,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "username": user})

}

// RegisterAdmin godoc
// @Summary Register a Full Access account (ADMIN).
// @Description registering a user from full access.
// @Tags Auth
// @Param Body body RegisterInput true "the body to register a FULL ACCESS account (ADMIN)"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register-admin [post]
func RegisterAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password
	u.FullAccess = true

	_, err := u.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var FA string

	if u.FullAccess == true {
		FA = "true"
	}

	user := map[string]string{

		"username":    input.Username,
		"email":       input.Email,
		"full_access": FA,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "username": user})

}

// UpdatePassword godoc
// @Summary Change Password.
// @Description change password.
// @Tags Auth
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param Body body ChangePasswordInput true "the body to update age merk category"
// @Success 200 {object} models.User
// @Router /change-password [patch]
func UpdatePassword(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := new(models.User)
	user, _ := c.Get("user")
	users := user.(models.User)

	hashed, _ := u.GeneratePassword(input.Password)

	result := db.Model(u).Where("id = ?", users.ID).UpdateColumn("password", hashed)

	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):

		default:
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "password has been changed!", "username": users.Username})
}
