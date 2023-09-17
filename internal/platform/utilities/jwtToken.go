package utilities

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const JWT_SECRET = "jwt_secret"

func CreateJWTToken(email string, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
