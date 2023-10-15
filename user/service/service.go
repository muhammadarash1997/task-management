package service

import (
	"github.com/muhammadarash1997/task-management/user/repository"
	"github.com/muhammadarash1997/task-management/user/model"
	"log"
	"runtime"
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
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request :%v", file, line, user)

	err := s.repository.CreateUser(user)
	if err != nil {
		log.Printf("Error :%v", err)
		return err
	}

	log.Printf("[%v] [%v] Response :%v", file, line, user)
	return nil
}

func (s *service) FindUsers() ([]model.User, error) {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request", file, line)

	users, err := s.repository.GetAllUser()
	if err != nil {
		log.Printf("Error :%v", err)
		return []model.User{}, err
	}

	log.Printf("[%v] [%v] Response :%v", file, line, users)
	return users, nil
}

func (s *service) FindUser(userId uint) (model.User, error) {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request :%v", file, line, userId)
	
	user, err := s.repository.GetUserById(userId)
	if err != nil {
		log.Printf("Error :%v", err)
		return model.User{}, err
	}

	log.Printf("[%v] [%v] Response :%v", file, line, user)
	return user, nil
}

func (s *service) UpdateUser(user model.User) error {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request :%v", file, line, user)

	err := s.repository.UpdateUser(user)
	if err != nil {
		log.Printf("Error :%v", err)
		return err
	}

	log.Printf("[%v] [%v] Response", file, line)
	return nil
}

func (s *service) DeleteUser(user model.User) error {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request :%v", file, line, user)

	err := s.repository.DeleteById(user.ID)
	if err != nil {
		log.Printf("Error :%v", err)
		return err
	}

	log.Printf("[%v] [%v] Response", file, line)
	return nil
}