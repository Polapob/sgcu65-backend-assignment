package controllers

import (
	"net/http"
	"sgcu65/models"

	"github.com/gin-gonic/gin"
)

type AddTeamDTO struct {
	Name string
}

func (db *DBController) AddTeam(c *gin.Context) {

	var addedTeam AddTeamDTO

	if err := c.ShouldBindJSON(&addedTeam); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	team := models.Team{
		Name: addedTeam.Name,
	}

	if err := db.database.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}

func (db *DBController) DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := db.database.First(&team, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.database.Delete(&team, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "successfully Delete Team"})
}

func (db *DBController) GetTeam(c *gin.Context) {
	var team models.Team
	id := c.Param("id")

	if err := db.database.Find(&team, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}
