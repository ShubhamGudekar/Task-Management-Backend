package service

import (
	"Task-Management-Backend/internal/dto"
	"Task-Management-Backend/internal/model"
	"Task-Management-Backend/internal/repository"

	"github.com/jinzhu/copier"
)

func CreateTask(taskRequest dto.TaskRequest) (*dto.TaskResponse, error) {
	var task model.Task
	if err := copier.Copy(&task, taskRequest); err != nil {
		return nil, err
	}

	createdTask, err := repository.CreateTask(&task)
	if err != nil {
		return nil, err
	}

	var taskResponse dto.TaskResponse
	if err := copier.Copy(&taskResponse, createdTask); err != nil {
		return nil, err
	}
	return &taskResponse, nil
}

func GetAllTasks() ([]dto.TaskResponse, error) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var taskResponse []dto.TaskResponse
	if err := copier.Copy(&taskResponse, tasks); err != nil {
		return nil, err
	}

	return taskResponse, nil
}
