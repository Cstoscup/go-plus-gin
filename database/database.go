package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-plus-gin/model"
)

var DB *gorm.DB

func Connect() {
	psqlInfo := "host=localhost port=5432 user=calliestoscup password='' dbname=album_list sslmode=disable"

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		db.AutoMigrate(&model.Album{})
		DB = db
		fmt.Println("Connection Opened to Database")
	}
}
