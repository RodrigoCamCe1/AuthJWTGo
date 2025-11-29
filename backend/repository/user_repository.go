package repository

import (
	"JWTfinalfinalFrontandback/database"
	"JWTfinalfinalFrontandback/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func FindByEmail(email string) (models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func FindAll() ([]models.User, error) {
	var users []models.User
	// Find(&users) busca todos los registros y los guarda en el slice
	err := database.DB.Find(&users).Error
	return users, err
}
