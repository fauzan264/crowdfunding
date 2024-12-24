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
		DbHost = os.Getenv("DB_HOST")
		DbPort = os.Getenv("DB_PORT")
		DbName = os.Getenv("DB_NAME")
	)

	fmt.Println(DbUser)
	fmt.Println(DbPassword)
	fmt.Println(DbHost)
	fmt.Println(DbPort)
	fmt.Println(DbName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Println(dsn)
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
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "ahmad264"
	// userInput.Email = "ahmad264@mail.com"
	// userInput.Occupation = "golang developer"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)

	// userRepository.Save(user)
	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()
}

// input dari user, ini isinya cuma struct input
// handler -> menangkap input user, mapping ke object go(struct) -> logic/controller
// service -> mengambil object mapping ke struct user/entity
// repository, save struct user ke db -> db