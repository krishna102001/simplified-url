package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/krishna102001/simplify-url/database"
	"github.com/krishna102001/simplify-url/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load the env file")
	}
	database.Init_DB()
}

func main() {
	router := gin.Default()

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Println("Failed to set proxies")
	}

	router.POST("/create-short-url", routes.CreateSimplifyUrl)

	router.GET("/get-short-url/:id", routes.GetSimplifyUrl)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("Empty PORT")
	}
	if err := router.Run(":" + PORT); err != nil {
		log.Println("Server is running on Port no. : ", PORT)
	}
}
