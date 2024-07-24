package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn, exists := os.LookupEnv("DATABASE_DSN")
	if !exists {
		dsn = "db.sqlite3" // default to sqlite
	}

	database, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
	return DB
}
