package controllers

import (
	"net/http"
	"os"
	"sgcu65/models"
	"sgcu65/services"

	"github.com/gin-gonic/gin"
)

type AddUserInterface struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Role      string `json:"role"`
	Position  string `json:"position"`
	Salary    uint   `json:"salary"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

func (db *DBController) AddUser(c *gin.Context) {
	var addedUser AddUserInterface
	if err := c.ShouldBindJSON(&addedUser); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bcryptSecret := os.Getenv("bcryptSecret")
	password, err := services.HashPassword(bcryptSecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user := &models.User{
		Email:     addedUser.Email,
		Firstname: addedUser.Firstname,
		Surname:   addedUser.Surname,
		Role:      addedUser.Role,
		Position:  addedUser.Position,
		Salary:    addedUser.Salary,
		Password:  password,
		Username:  addedUser.Username,
	}

	if err := db.database.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func (db *DBController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.database.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := db.database.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "successfully delete user"})
}

func (db *DBController) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.database.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
