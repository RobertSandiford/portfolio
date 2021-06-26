package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"github.com/shopspring/decimal"

	. "main/models"
)

// Initiase GORM
func initDb() {

	//// gorm postgresql
	dsn := "host=" + os.Getenv("postgres_host") + " user=" + os.Getenv("postgres_user") + " password=" + os.Getenv("postgres_pass") + " dbname=" + os.Getenv("postgres_db") + " port=" + os.Getenv("postgres_port") + " sslmode=disable TimeZone=Europe/London"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// setup database??
	db.AutoMigrate(&Main{})
	db.AutoMigrate(&Asset{})
	db.AutoMigrate(&Trade{})
	db.AutoMigrate(&Position{})

	main := Main{}
	result := db.First(&main)
	if result.RowsAffected == 0 {
		main.Pool, err = decimal.NewFromString("1")
		if err != nil {
			panic(err)
		}
	}

}
