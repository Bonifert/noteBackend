package service

import (
	"awesomeProject/database"
	"awesomeProject/dto"
	"awesomeProject/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(newUser *dto.NewUser) (uint, error) {
	db := database.DB
	password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user := model.User{
		Username: newUser.Username,
		Password: string(password),
	}
	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func getUserByName(username string) (model.User, error) {
	db := database.DB
	user := model.User{}
	result := db.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func GetUserById(userId uint) (model.User, error) {
	user := model.User{}
	result := database.DB.Where("ID = ?", userId).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}
