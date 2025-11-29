package utils

import (
	"JWTfinalfinalFrontandback/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string, role string) (string, error) {
	secret := []byte(config.GetEnv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"email": email,
		"role":  role, // Guardamos el rol en el token
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := []byte(config.GetEnv("JWT_SECRET"))
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
