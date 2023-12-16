package config

import (
	"Depublic-App-Service/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func DatabaseInit() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
	}

	db.AutoMigrate(&model.User{})
}

func GetDB() *gorm.DB {
	return db
}
