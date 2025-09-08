package controller

import (
	"Task-Management-Backend/internal/dto"
	user_errors "Task-Management-Backend/internal/errors"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Encrpyt Passowrd
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": userResponse})
}

func Login(c *gin.Context) {
	// Map Request Body
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get User details from database
	user, err := service.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid Email or Password"})
		return
	}

	// Validate Password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid Email or Password"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret not configured"})
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set Cookie
	// SameSite important to prevent against CRSF attacks
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, expiryTime, "", "", false, true)
	c.Status(http.StatusOK)
}

func ForgotPassword() {

}
