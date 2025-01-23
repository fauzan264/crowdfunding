package main

import (
	"github.com/fauzan264/crowdfunding/backend/auth"
	"github.com/fauzan264/crowdfunding/backend/campaign"
	"github.com/fauzan264/crowdfunding/backend/config"
	"github.com/fauzan264/crowdfunding/backend/handler"
	"github.com/fauzan264/crowdfunding/backend/middleware"
	"github.com/fauzan264/crowdfunding/backend/user"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDatabase()

	// Repositories
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	
	// Services
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	// Handlers
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// Initialize router
	router := gin.Default()
	router.Static("/images", "./images")
	
	api := router.Group("/api/v1")

	// User-related routes
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibility)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	// Campaign-related routes
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns", middleware.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(authService, userService), campaignHandler.UploadImage)

	router.Run()
}