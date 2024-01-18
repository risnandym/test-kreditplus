package main

import (
	"kredit_plus/core/config"
	"kredit_plus/core/utils"
	"kredit_plus/docs"
	"kredit_plus/src"
	"kredit_plus/src/routes"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// for load godotenv
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Review Mobile Phone API"
	docs.SwaggerInfo.Description = "This is a Final Project Golang JCC."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	err := config.Init()
	if err != nil {
		panic(err)
	}

	// // database connection
	// db,_ := config.ConnectDataBase(config.Postgres())
	// sqlDB, _ := db.DB()
	// defer sqlDB.Close()

	app := src.Dependencies()

	// router
	r := routes.SetupRouter(app)
	// just remove port 8080
	r.Run()
}
