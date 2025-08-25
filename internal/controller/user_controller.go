package controller

import (
	"Task-Management-Backend/internal/dto"
	user_errors "Task-Management-Backend/internal/errors"
	"Task-Management-Backend/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := service.CreateUser(&req)
	if err != nil {
		if errors.Is(err, user_errors.ErrEmailAlreadyRegistered) {
			c.JSON(http.StatusConflict, gin.H{"error": user_errors.ErrEmailAlreadyRegistered.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": userResponse})
}

func GetAllUsers(c *gin.Context) {

	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserById(c *gin.Context) {

	id := c.Param("id")

	user, err := service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUserById(c *gin.Context) {

	id := c.Param("id")
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.UpdateUser(id, &req)
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"user": user})
}

func DeleteUserById(c *gin.Context) {

	id := c.Param("id")

	user, err := service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
