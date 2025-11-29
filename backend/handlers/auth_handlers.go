package handlers

import (
	"JWTfinalfinalFrontandback/repository"
	"JWTfinalfinalFrontandback/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body map[string]string
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := services.Register(body["email"], body["password"], body["role"]); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado exitosamente"})
}

func Login(c *gin.Context) {
	var body map[string]string
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	token, err := services.Login(body["email"], body["password"])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func AdminEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "¡Hola Jefe! Bienvenido al panel de Administración."})
}

func GetAllUsers(c *gin.Context) {
	users, err := repository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}
