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

	// Auth Routes
	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.POST("/forgotPassword", controller.ForgotPassword)
	router.POST("/logout", middleware.ValidateAuthorization, controller.Logout)

	// User Routes
	router.GET("/users", controller.GetAllUsers)
	router.GET("/users/:id", controller.GetUserById)
	router.PUT("/users/:id", controller.UpdateUserById)
	router.DELETE("/users/:id", controller.DeleteUserById)
	router.GET("/users/:id/tasks", controller.GetAllTaskByUserId)

	// Task Routes
	router.POST("/tasks", controller.CreateTask)
	router.GET("/tasks", controller.GetAllTasks)
	router.GET("/tasks/:id", controller.GetTaskById)
	router.PUT("/tasks/:id", controller.UpdateTaskById)
	router.DELETE("/tasks/:id", controller.DeleteTaskById)

	router.Run()
}
