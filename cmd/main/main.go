package main

import (
	"kredit_plus/app"
	"kredit_plus/core/config"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {

	err := config.Init()
	if err != nil {
		panic(err)
	}

	// // database connection
	// db,_ := config.ConnectDataBase(config.Postgres())
	// sqlDB, _ := db.DB()
	// defer sqlDB.Close()

	dep := app.Dependencies()

	// router
	r := app.SetupRouter(dep)
	// just remove port 8080
	r.Run()
}
