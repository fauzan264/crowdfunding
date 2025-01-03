package main

import (
	"log"

	"github.com/fauzan264/crowdfunding/backend/auth"
	"github.com/fauzan264/crowdfunding/backend/config"
	"github.com/fauzan264/crowdfunding/backend/handler"
	"github.com/fauzan264/crowdfunding/backend/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDatabase()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibility)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}