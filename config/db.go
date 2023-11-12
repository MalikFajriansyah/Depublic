package config

import (
	"Depublic-App-Service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	dsn := "host=localhost user=postgres password=Malik1011* dbname=depublic port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	database.AutoMigrate(&model.User{}, &model.Event{})
}

func DB() *gorm.DB {
	return database
}
