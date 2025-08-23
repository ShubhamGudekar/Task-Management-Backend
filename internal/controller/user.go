package controller

import (
	"Task-Management-Backend/internal/infrastructure"
	"Task-Management-Backend/internal/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type userRequest struct {
	Name     string
	Email    string
	Password string
}

type userResponse struct {
	Id           int
	Name         string
	Email        string
	RegisteredAt time.Time `copier:"CreatedAt"`
	UpdatedAt    time.Time
}

func CreateUser(c *gin.Context) {

	// Bind incoming JSON to the struct
	var u userRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, err.Error())
		return
	}

	// Map request data to model
	var user model.User
	if err := copier.Copy(&user, &u); err != nil {
		c.JSON(500, err.Error())
	}

	// Store it in Database
	if result := infrastructure.DB.Create(&user); result.Error != nil {
		c.JSON(500, result.Error)
	}

	// Map result to response
	var ur userResponse
	if err := copier.Copy(&ur, &user); err != nil {
		c.JSON(500, err.Error())
	}

	// Return the response
	c.JSON(201, gin.H{"User": ur})

}

func GetAllUsers(c *gin.Context) {

	// Get All Users from Database
	var u []model.User
	if result := infrastructure.DB.Find(&u); result.Error != nil {
		c.JSON(500, result.Error)
	}

	// Map resukt to response
	var ur []userResponse
	if err := copier.Copy(&ur, &u); err != nil {
		c.JSON(500, err.Error())
	}

	//Return response
	c.JSON(200, ur)
}

func GetUserById(c *gin.Context) {

	// Parse ID from URL param
	id := c.Param("id")

	// Find User details
	var u model.User
	if result := infrastructure.DB.First(&u, id); result.Error != nil {
		c.JSON(500, result.Error)
	}

	//Map result to response
	var ur userResponse
	if err := copier.Copy(&ur, &u); err != nil {
		c.JSON(500, err.Error())
	}

	// Return Response
	c.JSON(200, ur)
}

func UpdateUserById(c *gin.Context) {

	// Parse ID from URL param
	id := c.Param("id")

	// Find User details
	var u model.User
	if result := infrastructure.DB.First(&u, id); result.Error != nil {
		c.JSON(500, result.Error)
	}

	// Bind incoming JSON to the struct
	var ud userRequest
	if err := c.ShouldBindJSON(&ud); err != nil {
		c.JSON(400, err.Error())
	}

	// Map request data to model
	if err := copier.Copy(&u, &ud); err != nil {
		c.JSON(500, err.Error())
	}

	// Store it in Database
	if result := infrastructure.DB.Save(&u); result.Error != nil {
		c.JSON(500, result.Error)
	}

	// Map result to response
	var ur userResponse
	if err := copier.Copy(&ur, &u); err != nil {
		c.JSON(500, err.Error())
	}

	// Return the response
	c.JSON(200, gin.H{"User": ur})
}

func DeleteUserById(c *gin.Context) {
	// Parse ID from URL param
	id := c.Param("id")

	// Find User details
	var u model.User
	if result := infrastructure.DB.First(&u, id); result.Error != nil {
		c.JSON(500, result.Error)
	}

	// Delete User
	if result := infrastructure.DB.Delete(&u); result.Error != nil {
		c.JSON(500, result.Error)
	}

	// Map result to response
	var ur userResponse
	if err := copier.Copy(&ur, &u); err != nil {
		c.JSON(500, err.Error())
	}

	// Return the response
	c.JSON(200, gin.H{"Deleted": ur})
}
