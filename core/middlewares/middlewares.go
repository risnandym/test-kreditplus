package middlewares

import (
	"log"
	"net/http"
	"test-kreditplus/core/config"
	"test-kreditplus/core/utils"
	"test-kreditplus/src/handlers"

	"github.com/gin-gonic/gin"
)

func PublicMiddleware(svc handlers.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := utils.TokenValid(c, config.Config())
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		id, _ := utils.ExtractTokenID(c, config.Config())

		auth, err := svc.Get(id)
		if err != nil {
			log.Fatal(err)
		}

		c.Set("auth", auth)

		c.Next()
	}
}
