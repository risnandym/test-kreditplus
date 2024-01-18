package middlewares

import (
	"kredit_plus/core/utils/token"
	"kredit_plus/src/app/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PublicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		id, _ := token.ExtractTokenID(c)
		db := c.MustGet("app").(*gorm.DB)

		var user entities.User
		result := db.First(&user, "id = ?", id)
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		c.Set("user", user)

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		id, _ := token.ExtractTokenID(c)
		db := c.MustGet("app").(*gorm.DB)

		var user entities.User
		result := db.First(&user, "id = ?", id)
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		// // Validate access
		// if user.FullAccess == false {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "you dont have an access!"})
		// 	return
		// }
		c.Set("user", user)

		c.Next()
	}
}
