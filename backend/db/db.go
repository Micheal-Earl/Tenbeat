package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mikesprogram.com/tenbeat/models"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
