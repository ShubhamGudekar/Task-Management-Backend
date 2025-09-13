package service

import (
	"Task-Management-Backend/internal/dto"
	"Task-Management-Backend/internal/errors"
	"Task-Management-Backend/internal/model"
	"Task-Management-Backend/internal/repository"

	"github.com/jinzhu/copier"
)

func CreateTask(userId uint, taskRequest dto.TaskRequest) (*dto.TaskResponse, error) {
	var task model.Task
	if err := copier.Copy(&task, taskRequest); err != nil {
		return nil, err
	}

	if task.Priority != "" && !task.IsValidPriority() {
		return nil, errors.ErrInvalidTaskPriority
	}

	if task.Status != "" && !task.IsValidStatus() {
		return nil, errors.ErrInvalidTaskStatus
	}

	task.UserID = userId

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

func GetAllTasksByUserId(id string) ([]dto.TaskResponse, error) {
	tasks, err := repository.GetAllTaskByUserId(id)
	if err != nil {
		return nil, err
	}

	var taskResponse []dto.TaskResponse
	if err := copier.Copy(&taskResponse, tasks); err != nil {
		return nil, err
	}

	return taskResponse, nil
}

func GetTaskById(taskId string, userId uint) (*dto.TaskResponse, error) {
	task, err := repository.GetTaskById(taskId, userId)
	if err != nil {
		return nil, err
	}

	var taskResponse dto.TaskResponse
	if err := copier.Copy(&taskResponse, task); err != nil {
		return nil, err
	}

	return &taskResponse, nil
}

func UpdateTask(taskId string, userId uint, taskRequest dto.TaskRequest) (*dto.TaskResponse, error) {
	task, err := repository.GetTaskById(taskId, userId)
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(&task, taskRequest); err != nil {
		return nil, err
	}

	if !task.IsValidPriority() {
		return nil, errors.ErrInvalidTaskPriority
	}

	if !task.IsValidStatus() {
		return nil, errors.ErrInvalidTaskStatus
	}

	task, err = repository.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	var taskResponse dto.TaskResponse
	if err = copier.Copy(&taskResponse, task); err != nil {
		return nil, err
	}

	return &taskResponse, nil
}

func DeleteTaskById(taskId string, userId uint) (*dto.TaskResponse, error) {
	task, err := repository.GetTaskById(taskId, userId)
	if err != nil {
		return nil, err
	}

	task, err = repository.DeleteTask(task)
	if err != nil {
		return nil, err
	}

	var taskResponse dto.TaskResponse
	if err = copier.Copy(&taskResponse, task); err != nil {
		return nil, err
	}

	return &taskResponse, nil
}
