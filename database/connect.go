package database

import (
	"awesomeProject/config"
	"awesomeProject/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	host := config.Config("HOST")
	user := config.Config("USER")
	password := config.Config("PASSWORD")
	dbname := config.Config("DBNAME")
	port := config.Config("PORT")
	sslmode := config.Config("SSLMODE")
	timezone := config.Config("TIMEZONE")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	err = db.AutoMigrate(&model.User{}, &model.Note{})
	if err != nil {
		return
	}

	DB = db
}
