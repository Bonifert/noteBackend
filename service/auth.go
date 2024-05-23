package service

import (
	"awesomeProject/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.RegisteredClaims
}

func Authenticate(username string, password string) (string, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	return generateJWT(username, user.ID)
}

func generateJWT(username string, id uint) (string, error) {
	claims :=
		JwtCustomClaims{
			Username: username,
			ID:       strconv.Itoa(int(id)),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			},
		}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := config.Config("SECRET_KEY")
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
