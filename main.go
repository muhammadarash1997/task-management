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

	// Routing untuk Auth
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authController.LoginHandler)
	}

	// Routing untuk Users
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userController.CreateUserHandler)
		userGroup.GET("/", userController.GetAllUsersHandler)
		userGroup.GET("/:id", userController.GetUserHandler)
		userGroup.PUT("/:id", userController.UpdateUserHandler)
		userGroup.DELETE("/:id", userController.DeleteUserHandler)
	}

	// Routing untuk Tasks yang dilindungi oleh OAuth2
	taskGroup := r.Group("/tasks")
	taskGroup.Use() // Middleware OAuth2 untuk melindungi rute
	{
		taskGroup.POST("/", authController.AuthenticateHandler, taskController.CreateTaskHandler)
		taskGroup.GET("/", authController.AuthenticateHandler, taskController.GetAllTasksHandler)
		taskGroup.GET("/:id", authController.AuthenticateHandler, taskController.GetTaskHandler)
		taskGroup.PUT("/:id", authController.AuthenticateHandler, taskController.UpdateTaskHandler)
		taskGroup.DELETE("/:id", authController.AuthenticateHandler, taskController.DeleteTaskHandler)
	}

	r.Run(":8080")
}
