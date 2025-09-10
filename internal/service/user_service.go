package service

import (
	"Task-Management-Backend/internal/dto"
	"Task-Management-Backend/internal/model"
	"Task-Management-Backend/internal/repository"

	"github.com/jinzhu/copier"
)

func CreateUser(userRequest *dto.SignUpRequest) (*dto.UserResponse, error) {
	var user model.User
	if err := copier.Copy(&user, userRequest); err != nil {
		return nil, err
	}

	createdUser, err := repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	var userResponse dto.UserResponse
	if err := copier.Copy(&userResponse, createdUser); err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func GetAllUsers() ([]dto.UserResponse, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userResponse []dto.UserResponse
	if err := copier.Copy(&userResponse, users); err != nil {
		return nil, err
	}

	return userResponse, nil
}

func GetUserByID(id string) (*dto.UserResponse, error) {

	user, err := repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	var userResponse dto.UserResponse
	if err := copier.Copy(&userResponse, user); err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func GetUserByEmail(email string) (*model.User, error) {

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id string, userRequest *dto.UserRequest) (*dto.UserResponse, error) {

	user, err := repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(&user, userRequest); err != nil {
		return nil, err
	}

	user, err = repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	var userResponse dto.UserResponse
	if err := copier.Copy(&userResponse, user); err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func DeleteUser(id string) (*dto.UserResponse, error) {

	user, err := repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	user, err = repository.DeleteUser(user)
	if err != nil {
		return nil, err
	}

	var userResponse dto.UserResponse
	if err := copier.Copy(&userResponse, user); err != nil {
		return nil, err
	}

	return &userResponse, nil
}
