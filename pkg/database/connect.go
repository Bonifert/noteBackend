package database

import (
	"awesomeProject/pkg/config"
	model2 "awesomeProject/pkg/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	host := config.Config("DB_HOST")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	dbname := config.Config("DB_NAME")
	port := config.Config("DB_PORT")
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

	err = db.AutoMigrate(&model2.User{}, &model2.Note{})
	if err != nil {
		return
	}

	DB = db
}
