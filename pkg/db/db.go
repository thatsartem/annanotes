package db

import (
	"annanotes/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Note{}, &models.User{})

	return db
}
