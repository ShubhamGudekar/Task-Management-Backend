package controller

import (
	"Task-Management-Backend/internal/dto"
	task_errors "Task-Management-Backend/internal/errors"
	"Task-Management-Backend/internal/middleware"
	"Task-Management-Backend/internal/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var taskRequest dto.TaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskResponse, err := service.CreateTask(user.ID, taskRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": taskResponse})
}

func GetAllTaskByUserId(c *gin.Context) {

	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tasks, err := service.GetAllTasksByUserId(fmt.Sprintf("%d", user.ID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTaskById(c *gin.Context) {

	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	taskId := c.Param("id")

	task, err := service.GetTaskById(taskId, user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func UpdateTaskById(c *gin.Context) {

	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	taskId := c.Param("id")
	var req dto.TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := service.UpdateTask(taskId, user.ID, req)
	if err != nil {
		if errors.Is(err, task_errors.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": task_errors.ErrTaskNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func DeleteTaskById(c *gin.Context) {

	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	taskId := c.Param("id")

	task, err := service.DeleteTaskById(taskId, user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})

}
