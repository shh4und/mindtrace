package sqlite

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}
