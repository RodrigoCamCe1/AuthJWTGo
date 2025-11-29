package main

import (
	"JWTfinalfinalFrontandback/database"
	"JWTfinalfinalFrontandback/handlers"
	"JWTfinalfinalFrontandback/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/api/auth/register", handlers.Register)
	r.POST("/api/auth/login", handlers.Login)

	// Rutas Protegidas
	protected := r.Group("/api")
	protected.Use(middlewares.JWTMiddleware())
	{
		protected.GET("/protected", handlers.ProtectedEndpoint)
	}

	// Rutas Admin
	admin := r.Group("/api/admin")
	admin.Use(middlewares.JWTMiddleware())
	admin.Use(middlewares.RoleMiddleware("ROLE_ADMIN"))
	{
		admin.GET("/dashboard", handlers.AdminEndpoint)
		admin.GET("/users", handlers.GetAllUsers)
	}

	r.Run(":8080")
}
