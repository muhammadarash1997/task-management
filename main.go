package main

import (
	"log"

	"github.com/gin-gonic/gin"
	authcontroller "github.com/muhammadarash1997/task-management/auth/controller"
	authservice "github.com/muhammadarash1997/task-management/auth/service"
	"github.com/muhammadarash1997/task-management/database"
	taskcontroller "github.com/muhammadarash1997/task-management/task/controller"
	taskrepository "github.com/muhammadarash1997/task-management/task/repository"
	taskservice "github.com/muhammadarash1997/task-management/task/service"
	usercontroller "github.com/muhammadarash1997/task-management/user/controller"
	userrepository "github.com/muhammadarash1997/task-management/user/repository"
	userservice "github.com/muhammadarash1997/task-management/user/service"
)

func main() {
	r := gin.Default()
	r.Use(authservice.CORSMiddleware())

	db, err := database.StartConnection()
	if err != nil {
		log.Printf("Error :%v", err)
		return
	}
	
	userRepository := userrepository.NewRepository(db)
	taskRepository := taskrepository.NewRepository(db)

	userService := userservice.NewService(userRepository)
	taskService := taskservice.NewService(taskRepository)
	authService := authservice.NewService(userRepository)
	
	userController := usercontroller.NewController(userService)
	taskController := taskcontroller.NewController(taskService)
	authController := authcontroller.NewController(authService, userService)

	// Routing untuk Tasks yang dilindungi oleh OAuth2
	userGroup := r.Group("/users")
	userGroup.Use() // Middleware OAuth2 untuk melindungi rute
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}

	// Routing untuk Tasks yang dilindungi oleh OAuth2
	taskGroup := r.Group("/tasks")
	taskGroup.Use() // Middleware OAuth2 untuk melindungi rute
	{
		taskGroup.POST("/", authController.AuthenticateHandler, taskController.CreateTask)
		taskGroup.GET("/", authController.AuthenticateHandler, taskController.GetAllTasks)
		taskGroup.GET("/:id", authController.AuthenticateHandler, taskController.GetTask)
		taskGroup.PUT("/:id", authController.AuthenticateHandler, taskController.UpdateTask)
		taskGroup.DELETE("/:id", authController.AuthenticateHandler, taskController.DeleteTask)
	}

	r.Run(":8080")
}
