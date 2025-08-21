package main

import (
	"Task-Management-Backend/internal/infrastructure"

	"github.com/gin-gonic/gin"
)

func init() {
	infrastructure.LoadEnvVariables()
	infrastructure.ConnectDatabase()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
