package database

import (
	"JWTfinalfinalFrontandback/config"
	"JWTfinalfinalFrontandback/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config.LoadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	// Auto-migración: Crea la tabla 'users' automáticamente
	DB.AutoMigrate(&models.User{})

	log.Println("🚀 Conectado a PostgreSQL exitosamente.")
}
