package repository

import (
	task_error "Task-Management-Backend/internal/errors"
	"Task-Management-Backend/internal/infrastructure"
	"Task-Management-Backend/internal/model"
	"errors"

	"gorm.io/gorm"
)

func CreateTask(task *model.Task) (*model.Task, error) {
	if err := infrastructure.DB.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	if err := infrastructure.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetAllTaskByUserId(id string) ([]model.Task, error) {
	var tasks []model.Task
	if err := infrastructure.DB.Where("UserId=?", id).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil

}

func GetTaskById(id string) (*model.Task, error) {
	var task model.Task
	if err := infrastructure.DB.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, task_error.ErrTaskNotFound
		}
	}
	return &task, nil
}

func UpdateTask(task *model.Task) (*model.Task, error) {
	if err := infrastructure.DB.Save(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func DeleteTask(task *model.Task) (*model.Task, error) {

	if err := infrastructure.DB.Delete(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}
