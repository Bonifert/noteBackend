package service

import (
	"awesomeProject/database"
	"awesomeProject/dto"
	"awesomeProject/model"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
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
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == "23505" {
				return 0, ErrDuplicated
			}
		}
		return 0, err
	}
	return user.ID, nil
}

func EditUsernameById(id uint, editUser *dto.NewUsername) error {
	user, err := GetUserById(id)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(editUser.Password)); err != nil {
		return ErrUnauthorized
	}
	user.Username = editUser.NewUsername
	err = database.DB.Save(&user).Error
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == "23505" {
				return ErrDuplicated
			}
		}
	}
	return nil
}

func EditPasswordById(id uint, editPassword dto.NewPassword) error {
	user, err := GetUserById(id)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(editPassword.Password)); err != nil {
		return ErrUnauthorized
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(newPassword)
	database.DB.Save(&user)
	return nil
}

func GetUserByUsername(username string) (*model.User, error) {
	db := database.DB
	user := model.User{}
	result := db.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		return nil, ErrNotFound
	}
	return &user, nil
}

func GetUserById(userId uint) (*model.User, error) {
	user := model.User{}
	result := database.DB.Where("ID = ?", userId).First(&user)
	if result.Error != nil {
		return nil, ErrNotFound
	}
	return &user, nil
}

func DeleteUserById(userId uint) error {
	user, err := GetUserById(userId)
	if err != nil {
		return err
	}
	result := database.DB.Delete(&user)
	if result.RowsAffected == 0 {
		return errors.New("user already deleted")
	}
	return nil
}
