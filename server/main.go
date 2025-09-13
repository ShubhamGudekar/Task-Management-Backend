package main

import (
	"Task-Management-Backend/internal/controller"
	"Task-Management-Backend/internal/infrastructure"
	"Task-Management-Backend/internal/middleware"
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
	router.GET("/users", middleware.ValidateAuthorization, controller.GetUserDetails)
	router.PUT("/users", middleware.ValidateAuthorization, controller.UpdateUser)
	router.DELETE("/users", middleware.ValidateAuthorization, controller.DeleteUser)

	// Task Routes
	router.GET("/tasks", middleware.ValidateAuthorization, controller.GetAllTaskByUserId)
	router.POST("/tasks", middleware.ValidateAuthorization, controller.CreateTask)
	router.GET("/tasks/:id", middleware.ValidateAuthorization, controller.GetTaskById)
	router.PUT("/tasks/:id", middleware.ValidateAuthorization, controller.UpdateTaskById)
	router.DELETE("/tasks/:id", middleware.ValidateAuthorization, controller.DeleteTaskById)

	router.Run()
}
