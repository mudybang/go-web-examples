package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mudybang/go-web-examples/goginapi/controller"
	"github.com/mudybang/go-web-examples/goginapi/database"
	"github.com/mudybang/go-web-examples/goginapi/middleware"
	"github.com/mudybang/go-web-examples/goginapi/model"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Post{})
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/post", controller.CreatePost)
	protectedRoutes.GET("/posts", controller.GetAllPosts)

	router.Run("localhost:8000")
	fmt.Println("Server running on port 8000")
}
