package controllers

import (
	"net/http"
	"sgcu65/models"
	"sgcu65/services"
	"time"

	"github.com/gin-gonic/gin"
)

type AddTaskDTO struct {
	Name     string    `json:"name" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	Status   bool      `json:"status"`
	Deadline time.Time `json:"deadline" binding:"required"`
}

type taskController struct {
	taskService services.TaskService
}

type TaskController interface {
	AddTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTask(c *gin.Context)
}

func NewTaskController(s services.TaskService) taskController {
	return taskController{
		taskService: s,
	}
}

func (u taskController) AddTask(c *gin.Context) {
	var taskAdded AddTaskDTO

	if err := c.ShouldBindJSON(&taskAdded); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var newTask = models.Task{
		Name:     taskAdded.Name,
		Content:  taskAdded.Content,
		Status:   taskAdded.Status,
		Deadline: taskAdded.Deadline,
	}

	task, err := u.taskService.AddTask(newTask)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (u taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := u.taskService.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := u.taskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Task successfully remove"})
}

func (u taskController) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := u.taskService.GetTask(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}
