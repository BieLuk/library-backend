package postgres

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(config *config.Config) error {
	logrus.Info("Initializing Postgres DB connection")
	var err error
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	logrus.Info("Postgres DB connection initialized successfully")
	return nil
}

func GetDB() *gorm.DB {
	return db
}
