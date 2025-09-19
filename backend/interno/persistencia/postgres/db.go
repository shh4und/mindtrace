package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db_USER := os.Getenv("DB_USER")
	db_PASS := os.Getenv("DB_PASSWORD")
	db_NAME := os.Getenv("DB_DB")

	DSN := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		db_USER,
		db_PASS,
		db_NAME,
	)

	return gorm.Open(postgres.Open(DSN), &gorm.Config{})
}
