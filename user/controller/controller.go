package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadarash1997/task-management/user/model"
	"github.com/muhammadarash1997/task-management/user/service"
)

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{
		service: service,
	}
}

func (c *controller) CreateUserHandler(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.service.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *controller) GetAllUsersHandler(ctx *gin.Context) {
	users, err := c.service.FindUsers()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *controller) GetUserHandler(ctx *gin.Context) {
	userIDString := ctx.Param("id")
	userID, _ := strconv.ParseUint(userIDString, 10, 64)
	user, err := c.service.FindUser(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *controller) UpdateUserHandler(ctx *gin.Context) {
	userIDString := ctx.Param("id")
	userID, _ := strconv.ParseUint(userIDString, 10, 64)
	user, err := c.service.FindUser(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateUser model.User
	err = ctx.ShouldBindJSON(&updateUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update informasi user
	user.Name = updateUser.Name
	user.Email = updateUser.Email

	err = c.service.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *controller) DeleteUserHandler(ctx *gin.Context) {
	userIDString := ctx.Param("id")
	userID, _ := strconv.ParseUint(userIDString, 10, 64)
	user, err := c.service.FindUser(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = c.service.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
