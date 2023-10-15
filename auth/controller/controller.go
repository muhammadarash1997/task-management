package controller

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muhammadarash1997/task-management/auth/model"
	"github.com/muhammadarash1997/task-management/auth/service"
	userservice "github.com/muhammadarash1997/task-management/user/service"
)

type controller struct {
	service     service.Service
	userService userservice.Service
}

func NewController(service service.Service, userService userservice.Service) *controller {
	return &controller{service, userService}
}

func (c *controller) LoginHandler(ctx *gin.Context) {
	var login model.LoginRequest

	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Login failed"})
		return
	}

	userLogged, err := c.service.Login(login)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Login failed"})
		return
	}

	tokenGenerated, err := c.service.GenerateToken(userLogged.ID, map[string]bool{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Login failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "Login succesfully",
		"token":   tokenGenerated,
	})
}

func (c *controller) AuthenticateHandler(ctx *gin.Context) {
	// Ambil token dari header
	tokenInput := ctx.GetHeader("Authorization")

	// Validasi apakah benar itu adalah bearer token
	if !strings.Contains(tokenInput, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	tokenWithoutBearer := strings.Split(tokenInput, " ")[1]

	// Validasi token apakah berlaku
	token, err := c.service.ValidateToken(tokenWithoutBearer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Mengubah token yang bertipe jwt.Token menjadi bertipe jwt.MapClaims
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	id := claim["user_id"].(uint)
	user, err := c.userService.FindUser(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	ctx.Set("currentUser", user)
}
