package main

import (
	"sgcu65/models"
	"sgcu65/router"
)

func main() {

	db := models.ConnectDatabase()

	router.SetUpRouter(db)

}
