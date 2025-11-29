package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "¡Acceso permitido! Eres un usuario autenticado."})
}
