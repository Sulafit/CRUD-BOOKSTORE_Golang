package db

import (
	// "os/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() *gorm.DB {

	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty"
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return DB
}
