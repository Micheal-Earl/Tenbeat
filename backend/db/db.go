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

	migrate(db, &models.Book{})
	migrate(db, &models.User{})
	migrate(db, &models.Post{})

	return db
}

func migrate(db *gorm.DB, model any) {
	err := db.AutoMigrate(&model)
	if err != nil {
		log.Fatalln(err)
	}
}
