package controller

import (
	"Task-Management-Backend/internal/dto"
	user_errors "Task-Management-Backend/internal/errors"
	"Task-Management-Backend/internal/middleware"
	"Task-Management-Backend/internal/service"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetUserDetails(c *gin.Context) {
	// Get User from content
	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Copy user data to response
	var userResponse dto.UserResponse
	if err := copier.Copy(&userResponse, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to map user data to response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userResponse})
}

func UpdateUser(c *gin.Context) {
	// Get User from content
	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Map Request Body
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := service.UpdateUser(fmt.Sprintf("%d", user.ID), &req)
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"user": updatedUser})
}

func DeleteUser(c *gin.Context) {
	// Get User from content
	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Delete user
	deletedUser, err := service.DeleteUser(fmt.Sprintf("%d", user.ID))
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": deletedUser})
}
