package repository

import (
	user_errors "Task-Management-Backend/internal/errors"
	"Task-Management-Backend/internal/infrastructure"
	"Task-Management-Backend/internal/model"
	"errors"
	"strings"

	"gorm.io/gorm"
)

func CreateUser(user *model.User) (*model.User, error) {

	if err := infrastructure.DB.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, user_errors.ErrEmailAlreadyRegistered
		}
	}
	return user, nil
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := infrastructure.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := infrastructure.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_errors.ErrUserNotFound
		}
	}
	return &user, nil
}

func UpdateUser(user *model.User) (*model.User, error) {

	if err := infrastructure.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(user *model.User) (*model.User, error) {

	if err := infrastructure.DB.Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := infrastructure.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, user_errors.ErrUserNotFound
	}
	return &user, nil
}
