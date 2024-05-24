package service

import (
	"awesomeProject/database"
	"awesomeProject/dto"
	"awesomeProject/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(newUser *dto.UsernameAndPassword) (uint, error) {
	db := database.DB
	password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user := model.User{
		Username: newUser.Username,
		Password: string(password),
	}
	err = db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func GetUserByUsername(username string) (model.User, error) {
	db := database.DB
	user := model.User{}
	result := db.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func GetUserById(userId uint) (model.User, error) {
	user := model.User{}
	result := database.DB.Where("ID = ?", userId).First(&user)
	if result.Error != nil {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func DeleteUserById(userId uint) error {
	user := model.User{}
	result := database.DB.Where("ID = ?", userId).Delete(&user)
	if result.Error != nil {
		return errors.New("unexpected error occurred")
	}
	if result.RowsAffected == 0 {
		return errors.New("user already deleted")
	}
	return nil
}
