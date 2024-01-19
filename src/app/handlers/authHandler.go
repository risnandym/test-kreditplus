package handlers

import (
	"net/http"
	"test-kreditplus/src/app/contract"

	"github.com/gin-gonic/gin"
)

type ChangePasswordInput struct {
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"phone"`
	Password string `json:"password" binding:"required"`
}

// LoginUser godoc
//
//	@Summary		Login.
//	@Description	Logging in to get jwt token to access admin or user api by roles.
//	@Tags			Auth
//	@Param			Body	body	contract.LoginInput	true	"the body to login a user"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/kredit-plus/customer/login [post]
func Login(svc AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.ValidateAndBuildUserLogin(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Login(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "login success", "token": response})
	}
}

// Register godoc
//
//	@Summary		Register a User.
//	@Description	registering a user from public access.
//	@Tags			Auth
//	@Param			Body	body	contract.RegisterInput	true	"the body to register a user"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/kredit-plus/customer/register [post]
func Register(svc AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.ValidateAndBuildUserRegister(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Create(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "registration success", "data": response})
	}
}

// RegisterAdmin godoc
//	@Summary		Register a Full Access account (ADMIN).
//	@Description	registering a user from full access.
//	@Tags			Auth
//	@Param			Body	body	RegisterInput	true	"the body to register a FULL ACCESS account (ADMIN)"
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/register-admin [post]
// func RegisterAdmin(c *gin.Context) {
// 	db := c.MustGet("app").(*gorm.DB)
// 	var input RegisterInput

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	u := entities.User{}

// 	u.Email = input.Email
// 	u.Email = input.Email
// 	u.Password = input.Password
// 	u.LastLogin = time.Now()

// 	_, err := u.SaveUser(db)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := map[string]string{

// 		"username":   input.Email,
// 		"email":      input.Email,
// 		"last_login": u.LastLogin.Format("2006-01-02 15:04:05"),
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "registration success", "username": user})

// }

// UpdatePassword godoc
//	@Summary		Change Password.
//	@Description	change password.
//	@Tags			Auth
//	@Param			Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security		BearerToken
//	@Produce		json
//	@Param			Body	body		ChangePasswordInput	true	"the body to update age merk category"
//	@Success		200		{object}	entities.User
//	@Router			/change-password [patch]
// func UpdatePassword(c *gin.Context) {

// 	db := c.MustGet("app").(*gorm.DB)
// 	// Validate input
// 	var input ChangePasswordInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	u := new(entities.User)
// 	user, _ := c.Get("user")
// 	users := user.(entities.User)

// 	hashed, _ := u.GeneratePassword(input.Password)

// 	result := db.Model(u).Where("id = ?", users.ID).UpdateColumn("password", hashed)

// 	if result.Error != nil {
// 		switch {
// 		case errors.Is(result.Error, gorm.ErrRecordNotFound):

// 		default:
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "password has been changed!", "username": users.Email})
// }
