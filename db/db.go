package db

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Fatal("Failed to get database connection", zap.Error(err))
		return nil, fmt.Errorf("failed to get database connection: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		zap.L().Error("Failed to ping database", zap.Error(err))
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}
	return db, nil
}
