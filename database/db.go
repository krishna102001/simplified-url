package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init_DB() {
	dsn := os.Getenv("POSTGRESDBURL")
	if dsn == "" {
		log.Println("not found db url")
	}

	postgresdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to the database!", err)
	}
	DB = postgresdb
	log.Println("Successfully connected to database")
	MigrateTables()
}

func MigrateTables() {
	err := DB.AutoMigrate(&Shorturl{})
	if err != nil {
		log.Fatal("Failed to migrate the table")
	}
}
