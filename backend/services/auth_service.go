package services

import (
	"JWTfinalfinalFrontandback/models"
	"JWTfinalfinalFrontandback/repository"
	"JWTfinalfinalFrontandback/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Register(email, password string, role string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	// Si no envían rol, asignamos ROLE_USER por defecto
	if role == "" {
		role = "ROLE_USER"
	}

	user := models.User{
		Email:    email,
		Password: string(hashed),
		Role:     role, // Usamos el rol que nos pasaron
	}

	return repository.CreateUser(&user)
}

func Login(email, password string) (string, error) {
	user, err := repository.FindByEmail(email)
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	// Comparar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("contraseña incorrecta")
	}

	// Generar Token
	token, err := utils.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return "", errors.New("error generando token")
	}

	return token, nil
}
