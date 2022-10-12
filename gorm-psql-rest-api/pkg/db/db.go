package db

import (
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dbURL := "postgres://developer:developer@localhost:5432/rest_api"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
