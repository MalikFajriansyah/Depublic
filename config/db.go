package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseInit() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Malik1011* dbname=depublic port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		return nil, e
	}
	return database, nil
}
