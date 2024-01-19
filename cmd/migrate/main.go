package main

import (
	"test-kreditplus/core/config"
	"test-kreditplus/src/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db, _ := ConnectDataBase(cfg.Postgres)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

}

func ConnectDataBase(conn config.Postgres) (*gorm.DB, error) {

	username := conn.UserName
	password := conn.Password
	host := conn.Host
	port := conn.Port
	database := conn.Database
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.Auth{}, &entities.Profile{}, &entities.Limit{})

	return db, err
}
