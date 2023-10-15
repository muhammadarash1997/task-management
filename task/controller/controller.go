package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadarash1997/task-management/task/model"
	"github.com/muhammadarash1997/task-management/task/service"
)

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{
		service: service,
	}
}

func (c *controller) CreateTask(ctx *gin.Context) {
	var task model.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.service.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

func (c *controller) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.service.FindTasks()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *controller) GetTask(ctx *gin.Context) {
	taskIDString := ctx.Param("id")
	taskID, _ := strconv.ParseUint(taskIDString, 10, 64)
	task, err := c.service.FindTask(uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *controller) UpdateTask(ctx *gin.Context) {
	taskIDString := ctx.Param("id")
	taskID, _ := strconv.ParseUint(taskIDString, 10, 64)
	task, err := c.service.FindTask(uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updateTask model.Task
	err = ctx.ShouldBindJSON(&updateTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update informasi task
	task.Title = updateTask.Title
	task.Description = updateTask.Description
	task.Status = updateTask.Status

	err = c.service.UpdateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *controller) DeleteTask(ctx *gin.Context) {
	taskIDString := ctx.Param("id")
	taskID, _ := strconv.ParseUint(taskIDString, 10, 64)
	task, err := c.service.FindTask(uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	err = c.service.DeleteTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
