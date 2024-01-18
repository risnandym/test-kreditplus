package middlewares

import (
	"kredit_plus/app/src/handlers"
	"kredit_plus/core/config"
	"kredit_plus/core/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublicMiddleware(svc handlers.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("sudah masuk handler mdlwr")

		err := utils.TokenValid(c, config.Config())
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		log.Println("sudah lewat token valid")

		id, _ := utils.ExtractTokenID(c, config.Config())
		log.Println("sudah lewat token extract")

		auth, err := svc.Get(id)
		if err != nil {
			log.Fatal(err)
		}

		c.Set("auth", auth)

		c.Next()
	}
}
