package service

import (
	"github.com/muhammadarash1997/task-management/sharevar"
	"github.com/muhammadarash1997/task-management/user/model"
	"github.com/muhammadarash1997/task-management/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(model.User) error
	FindUsers() ([]model.User, error)
	FindUser(uint) (model.User, error)
	UpdateUser(model.User) error
	DeleteUser(model.User) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateUser(user model.User) error {
	sharevar.InfoLogger.Println("Request", user)

	//Hashing password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return err
	}
	user.Password = string(passwordHash)

	err = s.repository.CreateUser(user)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return err
	}

	sharevar.InfoLogger.Println("Response")
	return nil
}

func (s *service) FindUsers() ([]model.User, error) {
	sharevar.InfoLogger.Println("Request")

	users, err := s.repository.GetAllUser()
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return []model.User{}, err
	}

	sharevar.InfoLogger.Println("Response", users)
	return users, nil
}

func (s *service) FindUser(userId uint) (model.User, error) {
	sharevar.InfoLogger.Println("Request", userId)

	user, err := s.repository.GetUserById(userId)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return model.User{}, err
	}

	sharevar.InfoLogger.Println("Response", user)
	return user, nil
}

func (s *service) UpdateUser(user model.User) error {
	sharevar.InfoLogger.Println("Request", user)

	err := s.repository.UpdateUser(user)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return err
	}

	sharevar.InfoLogger.Println("Response")
	return nil
}

func (s *service) DeleteUser(user model.User) error {
	sharevar.InfoLogger.Println("Request", user)

	err := s.repository.DeleteById(user.ID)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return err
	}

	sharevar.InfoLogger.Println("Response")
	return nil
}
