package middlewares

import (
	"JWTfinalfinalFrontandback/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Obtener el token (Ya sabemos que existe gracias al JWTMiddleware previo)
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// 2. Parsear el token para leer los datos (Claims)
		token, _ := utils.ValidateToken(tokenString)
		claims, ok := token.Claims.(jwt.MapClaims)

		// 3. Verificar el Rol
		if ok && token.Valid {
			role := claims["role"].(string)
			if role != requiredRole {
				c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos de Administrador para estar aquí"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
