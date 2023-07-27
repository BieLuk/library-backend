package db

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Init(config config.Config) {
	var err error
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")
}

func GetDB() *gorm.DB {
	return db
}
