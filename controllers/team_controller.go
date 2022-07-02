package controllers

import (
	"net/http"
	"sgcu65/models"
	"sgcu65/services"

	"github.com/gin-gonic/gin"
)

type AddTeamDTO struct {
	Name string
}

type teamController struct {
	teamService services.TeamService
}

type TeamController interface {
	AddTeam(c *gin.Context)
	DeleteTeam(c *gin.Context)
	GetTeam(c *gin.Context)
}

func NewTeamController(s services.TeamService) TeamController {
	return teamController{
		teamService: s,
	}
}

func (t teamController) AddTeam(c *gin.Context) {

	var addedTeam AddTeamDTO

	if err := c.ShouldBindJSON(&addedTeam); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newTeam := models.Team{
		Name: addedTeam.Name,
	}

	team, err := t.teamService.AddTeam(newTeam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}

func (t teamController) DeleteTeam(c *gin.Context) {
	id := c.Param("id")

	_, err := t.teamService.GetTeam(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := t.teamService.DeleteTeam(id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "successfully Delete Team"})
}

func (t teamController) GetTeam(c *gin.Context) {
	id := c.Param("id")

	team, err := t.teamService.GetTeam(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}
