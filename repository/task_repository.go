package repository

import (
	"sgcu65/models"

	"gorm.io/gorm"
)

type taskRepository struct {
	DB *gorm.DB
}

type TaskRepository interface {
	AddTask(task models.Task) (models.Task, error)
	GetTask(taskId string) (models.Task, error)
	DeleteTask(taskId string) error
	Migrate() error
}

func NewtaskRepository(db *gorm.DB) TaskRepository {
	return taskRepository{
		DB: db,
	}
}

func (u taskRepository) Migrate() error {
	return u.DB.AutoMigrate(&models.Task{})
}

func (u taskRepository) AddTask(Task models.Task) (models.Task, error) {
	err := u.DB.Create(&Task).Error
	return Task, err
}

func (u taskRepository) GetTask(TaskId string) (models.Task, error) {
	var Task models.Task
	err := u.DB.Where("id = ?", TaskId).First(&Task).Error
	return Task, err
}

func (u taskRepository) DeleteTask(TaskId string) error {
	err := u.DB.Delete(&models.Task{}, TaskId).Error
	return err
}
