package controller

import (
	"Task-Management-Backend/internal/dto"
	user_errors "Task-Management-Backend/internal/errors"
	"Task-Management-Backend/internal/repository"
	"Task-Management-Backend/internal/service"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Map Request Body
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Encrpyt Passowrd
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process password"})
		return
	}
	req.Password = string(hashedPassword)

	// Create User
	userResponse, err := service.CreateUser(&req)
	if err != nil {
		if errors.Is(err, user_errors.ErrEmailAlreadyRegistered) {
			c.JSON(http.StatusConflict, gin.H{"error": user_errors.ErrEmailAlreadyRegistered.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": userResponse})
}

func Login(c *gin.Context) {
	// Map Request Body
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Get User details from database
	user, err := service.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid email or password"})
		return
	}

	// Validate Password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid email or password"})
		return
	}

	// Create a new token object
	var expiryTime int = 15 * 60
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     fmt.Sprintf("%v", user.ID),
		"expiry": time.Now().Add(time.Duration(expiryTime) * time.Second).Unix(),
	})

	// Get Secret from env variables
	secret := os.Getenv("JWTSECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt secret not configured"})
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set Cookie
	c.SetCookie("Auth", tokenString, expiryTime, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "user login successful"})
}

func ForgotPassword(c *gin.Context) {
	// Map Request Body
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Get User details from database
	user, err := service.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound.Error()})
		return
	}

	//Compare if old password is re-entered
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "new password can be same as old password"})
		return
	}

	// Encryt New Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to process password")
		return
	}

	// Save new password
	user.Password = string(hashedPassword)
	_, err = repository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to update password")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}

func Logout(c *gin.Context) {
	_, isExist := c.Get("user")
	if !isExist {
		c.JSON(http.StatusNotFound, gin.H{"error": user_errors.ErrUserNotFound.Error()})
		return
	}

	// Delete the JWT cookie by setting its MaxAge to -1
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", "", -1, "", "", false, true)

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "user logged out successfully"})
}
