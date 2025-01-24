package main

import (
	"log"

	"github.com/NyanLinHtoo/Oauth_Learning/config"
	"github.com/NyanLinHtoo/Oauth_Learning/handler"
	"github.com/NyanLinHtoo/Oauth_Learning/routes"
	"github.com/NyanLinHtoo/Oauth_Learning/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery())

	authService := services.NewAuthService()
	authHandler := handler.NewAuthHandler(authService)

	router := routes.NewRouter(authHandler)

	router.RegisterRoutes(r)

	log.Println("Server is running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
