package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muhammadarash1997/task-management/auth/model"
	"github.com/muhammadarash1997/task-management/sharevar"
	usermodel "github.com/muhammadarash1997/task-management/user/model"
	userrepository "github.com/muhammadarash1997/task-management/user/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	secretKey         string = "S3C123TKEY"
	tokenHourLifespan int64  = 24
)

type Service interface {
	Login(model.LoginRequest) (usermodel.User, error)
	GenerateToken(uint) (string, error)
	ValidateToken(string) (*jwt.Token, error)
}

type service struct {
	userRepository userrepository.Repository
}

func NewService(userRepository userrepository.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) Login(login model.LoginRequest) (usermodel.User, error) {
	sharevar.InfoLogger.Println("Request", login)

	email := login.Email
	password := login.Password

	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return user, err
	}

	sharevar.InfoLogger.Println("Response", user)
	return user, nil
}

func (s *service) GenerateToken(userID uint) (string, error) {
	sharevar.InfoLogger.Println("Request", userID)

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = time.Now().Add(time.Hour * time.Duration(tokenHourLifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenGenerated, err := token.SignedString([]byte(secretKey))
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return tokenGenerated, err
	}

	sharevar.InfoLogger.Println("Response", tokenGenerated)
	return tokenGenerated, nil
}

func (s *service) ValidateToken(encodedToken string) (*jwt.Token, error) {
	sharevar.InfoLogger.Println("Request", encodedToken)

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token")
		}

		return secretKey, nil
	})

	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return token, err
	}

	sharevar.InfoLogger.Println("Response", token)
	return token, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
