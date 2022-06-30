package main

import (
	"log"
	"sgcu65/controllers"
	"sgcu65/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := controllers.NewDBController(models.ConnectDatabase()) // new

	r.POST("/task/", db.AddTask)
	r.DELETE("/task/:id", db.DeleteTask)
	r.GET("/task/:id", db.GetTask)
	r.POST("/user/", db.AddUser)
	r.GET("/user/:id", db.GetUser)
	r.DELETE("/user/:id", db.DeleteUser)
	r.POST("/team", db.AddTeam)
	r.DELETE("/team/:id", db.DeleteTeam)
	r.GET("/team/:id", db.GetTeam)
	r.Run()
}
