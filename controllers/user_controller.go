package controllers

import (
	"net/http"
	"sgcu65/models"
	"sgcu65/services"

	"github.com/gin-gonic/gin"
)

type AddUserDTO struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Role      string `json:"role"`
	Position  string `json:"position"`
	Salary    uint   `json:"salary"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

type userController struct {
	userService services.UserService
}

type UserController interface {
	AddUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUser(c *gin.Context)
}

func NewUserController(s services.UserService) UserController {
	return userController{
		userService: s,
	}
}

func (u userController) AddUser(c *gin.Context) {
	var addedUser AddUserDTO
	if err := c.ShouldBindJSON(&addedUser); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	password, err := services.HashPassword(addedUser.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	addUser := models.User{
		Email:     addedUser.Email,
		Firstname: addedUser.Firstname,
		Surname:   addedUser.Surname,
		Role:      addedUser.Role,
		Position:  addedUser.Position,
		Salary:    addedUser.Salary,
		Password:  password,
		Username:  addedUser.Username,
	}

	user, err := u.userService.AddUser(addUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func (u userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := u.userService.GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := u.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "successfully delete user"})
}

func (u userController) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := u.userService.GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
