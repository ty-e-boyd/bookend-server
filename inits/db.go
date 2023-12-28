package inits

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func DBInit() {
	dsn := os.Getenv("POSTGRES_DEV_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Cannot connect to database")
	}

	DB = db
}
