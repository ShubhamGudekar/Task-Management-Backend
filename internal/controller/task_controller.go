package controller

import (
	"Task-Management-Backend/internal/dto"
	task_errors "Task-Management-Backend/internal/errors"
	"Task-Management-Backend/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var taskRequest dto.TaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskResponse, err := service.CreateTask(taskRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Task": taskResponse})
}

func GetAllTasks(c *gin.Context) {
	taskResponse, err := service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": taskResponse})
}

func GetAllTaskByUserId(c *gin.Context) {

	id := c.Param("id")

	tasks, err := service.GetAllTasksByUserId(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTaskById(c *gin.Context) {

	id := c.Param("id")

	task, err := service.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func UpdateTaskById(c *gin.Context) {

	id := c.Param("id")
	var req dto.TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := service.UpdateTask(id, req)
	if err != nil {
		if errors.Is(err, task_errors.ErrTaskNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": task_errors.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"task": task})
}

func DeleteTaskById(c *gin.Context) {

	id := c.Param("id")

	task, err := service.DeleteTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": task})

}
