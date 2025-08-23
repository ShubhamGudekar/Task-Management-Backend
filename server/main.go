package main

import (
	"Task-Management-Backend/internal/controller"
	"Task-Management-Backend/internal/infrastructure"
	"Task-Management-Backend/migrations"

	"github.com/gin-gonic/gin"
)

func init() {
	infrastructure.LoadEnvVariables()
	infrastructure.ConnectDatabase()
	migrations.MigrateDB()
}

func main() {
	router := gin.Default()

	// User Routes
	router.POST("/users", controller.CreateUser)
	router.GET("/users", controller.GetAllUsers)
	router.GET("/users/:id", controller.GetUserById)
	router.PUT("/users/:id", controller.UpdateUserById)
	router.DELETE("/users/:id", controller.DeleteUserById)

	router.Run()
}
