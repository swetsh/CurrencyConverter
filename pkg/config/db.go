package config

import (
	"CurrencyConverterService/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbname := "walletdb"
	dbuser := "postgres"
	password := "0000"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbuser,
		dbname,
		password,
	)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	d.AutoMigrate(&models.Currency{})

	if err != nil {
		log.Fatal(err)
	}

	db = d

	fmt.Println("Database connection successful...")
}

func GetDB() *gorm.DB {
	return db
}
