package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Holds our database connection
// configuration values
type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitializeDatabase(cfg DbConfig) (*gorm.DB, error) {
	// Create a new database connection
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.DBName, cfg.Password, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate Todo schema
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		return nil, err
	}

	// Migrate User schema
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}

	return db, nil

}
