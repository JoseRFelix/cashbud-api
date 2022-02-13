package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Account struct {
	gorm.Model
	ID      int     `json:"id" gorm:"primary_key"`
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
	Owner   string  `json:"owner"`
}

func ConnectDatabase() *gorm.DB {
	user := GetEnvironmentVariable("POSTGRES_USER")
	password := GetEnvironmentVariable("POSTGRES_PASSWORD")
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=cashbud port=5432", user, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	return db
}
