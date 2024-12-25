package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fauzan264/crowdfunding/backend/handler"
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
		DbUser = os.Getenv("DB_USER")
		DbPassword = os.Getenv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword)
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
	
	router.Run()
}

// input dari user, ini isinya cuma struct input
// handler -> menangkap input user, mapping ke object go(struct) -> logic/controller
// service -> mengambil object mapping ke struct user/entity
// repository, save struct user ke db -> db