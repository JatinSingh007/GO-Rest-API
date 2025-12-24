package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (error, int64) {
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// type checking the token
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("could not parse the token"), 0
	}

	if !parsedtoken.Valid {
		return errors.New("Invalid token!"), 0
	}
	// checking if mapclaims type
	claims, ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("Invalid token!"), 0
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return nil, userId

}
