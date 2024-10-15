package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

type DBConfig struct {
	gormDB *gorm.DB
}

// NewDB creates a new database connection
func NewDB() (*DBConfig, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(2 * time.Hour)

	return &DBConfig{gormDB: db}, nil
}

func (dbc *DBConfig) Close() error {
	sqlDB, err := dbc.gormDB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
