package controllers

import (
	"net/http"
	"sgcu65/models"
	"time"

	"github.com/gin-gonic/gin"
)

type AddTaskInterface struct {
	Name     string    `json:"name" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	Status   bool      `json:"status"`
	Deadline time.Time `json:"deadline" binding:"required"`
}

func (db *DBController) AddTask(c *gin.Context) {
	var taskAdded AddTaskInterface

	if err := c.ShouldBindJSON(&taskAdded); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var newTask = &models.Task{
		Name:     taskAdded.Name,
		Content:  taskAdded.Content,
		Status:   taskAdded.Status,
		Deadline: taskAdded.Deadline,
	}

	if err := db.database.Create(&newTask).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": taskAdded})
}

func (db *DBController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if err := db.database.First(&task, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := db.database.Delete(&task, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Task successfully remove"})
}

func (db *DBController) GetTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if err := db.database.First(&task, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}
