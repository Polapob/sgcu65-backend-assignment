package router

import (
	"log"
	"sgcu65/controllers"
	"sgcu65/repository"
	"sgcu65/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) {

	httpRouter := gin.Default()

	userRepository := repository.NewUserRepository(db)
	teamRepository := repository.NewTeamRepository(db)
	taskRepository := repository.NewtaskRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate error", err)
	}

	if err := teamRepository.Migrate(); err != nil {
		log.Fatal("User migrate error", err)
	}

	if err := taskRepository.Migrate(); err != nil {
		log.Fatal("User migrate error", err)
	}

	userService := services.NewUserService(userRepository)
	teamService := services.NewTeamService(teamRepository)
	taskService := services.NewTaskService(taskRepository)

	userController := controllers.NewUserController(userService)
	teamController := controllers.NewTeamController(teamService)
	taskController := controllers.NewTaskController(taskService)

	httpRouter.POST("/task/", taskController.AddTask)
	httpRouter.DELETE("/task/:id", taskController.DeleteTask)
	httpRouter.GET("/task/:id", taskController.GetTask)
	httpRouter.POST("/user/", userController.AddUser)
	httpRouter.GET("/user/:id", userController.GetUser)
	httpRouter.DELETE("/user/:id", userController.DeleteUser)
	httpRouter.POST("/team", teamController.AddTeam)
	httpRouter.DELETE("/team/:id", teamController.DeleteTeam)
	httpRouter.GET("/team/:id", teamController.GetTeam)
	httpRouter.Run()
}
