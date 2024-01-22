package main

import (
	"encoding/json"
	"fmt"
	"os"
	"test-kreditplus/core/config"
	"test-kreditplus/src"
	"test-kreditplus/src/contract"
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

	if err := db.Migrator().DropTable(&entities.Auth{}, &entities.Profile{}, &entities.Limit{}, &entities.Credit{}, &entities.Debit{}, &entities.Asset{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&entities.Auth{}, &entities.Profile{}, &entities.Limit{}, &entities.Credit{}, &entities.Debit{}, &entities.Asset{}); err != nil {
		return nil, err
	}

	// fmt.Print(SeedingtoDataBase().Error())

	return db, err
}

func SeedingtoDataBase() error {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	app := src.Dependencies()

	file, err := os.Open("./cmd/migrate/auth.json")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer file.Close()

	// Menggunakan json.NewDecoder untuk membaca data dari file JSON
	decoder := json.NewDecoder(file)

	// Membaca data JSON ke dalam slice RegisterInput
	var data []contract.RegisterInput
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Menampilkan data yang telah dibaca
	fmt.Println("Data yang dibaca dari file JSON:")
	for _, auth := range data {
		fmt.Printf("Username: %s, Password: %s\n", auth.Email, auth.Password)

		app.Services.AuthSVC.Create(auth)

	}

	return err
}
