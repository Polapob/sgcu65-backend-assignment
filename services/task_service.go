package services

import (
	"sgcu65/models"
	"sgcu65/repository"
)

type taskService struct {
	taskRepository repository.TaskRepository
}

type TaskService interface {
	AddTask(task models.Task) (models.Task, error)
	GetTask(taskId string) (models.Task, error)
	DeleteTask(taskId string) error
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return taskService{
		taskRepository: r,
	}
}

func (t taskService) AddTask(task models.Task) (models.Task, error) {
	return t.taskRepository.AddTask(task)
}

func (t taskService) GetTask(taskId string) (models.Task, error) {
	return t.taskRepository.GetTask(taskId)
}

func (t taskService) DeleteTask(taskId string) error {
	return t.taskRepository.DeleteTask(taskId)
}
