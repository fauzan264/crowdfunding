package main

import (
	"fmt"
	"log"

	"github.com/fauzan264/crowdfunding/backend/handler"
	"github.com/fauzan264/crowdfunding/backend/helper"
	"github.com/fauzan264/crowdfunding/backend/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var(
		DbHost = helper.GetEnv("DB_HOST", "db")
		DBPort = helper. GetEnv("DB_PORT", "3306")
		DbUser = helper.GetEnv("DB_USER", "root")
		DbPassword = helper.GetEnv("DB_PASSWORD", "root")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibility)

	router.Run()
}