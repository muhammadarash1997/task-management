package service

import (
	"github.com/muhammadarash1997/task-management/sharevar"
	"github.com/muhammadarash1997/task-management/task/model"
	"github.com/muhammadarash1997/task-management/task/repository"
)

type Service interface {
	CreateTask(model.Task) error
	FindTasks() ([]model.Task, error)
	FindTask(uint) (model.Task, error)
	UpdateTask(model.Task) error
	DeleteTask(model.Task) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateTask(task model.Task) error {
	sharevar.InfoLogger.Println("Request", task)

	err := s.repository.CreateTask(task)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return err
	}

	sharevar.InfoLogger.Println("Response")
	return nil
}

func (s *service) FindTasks() ([]model.Task, error) {
	sharevar.InfoLogger.Println("Request")

	tasks, err := s.repository.GetAllTask()
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return []model.Task{}, err
	}

	sharevar.InfoLogger.Println("Response", tasks)
	return tasks, nil
}

func (s *service) FindTask(taskId uint) (model.Task, error) {
	sharevar.InfoLogger.Println("Request", taskId)

	task, err := s.repository.GetTaskById(taskId)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return model.Task{}, err
	}

	sharevar.InfoLogger.Println("Response", task)
	return task, nil
}

func (s *service) UpdateTask(task model.Task) error {
	sharevar.InfoLogger.Println("Request", task)

	err := s.repository.UpdateTask(task)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return err
	}

	sharevar.InfoLogger.Println("Response")
	return nil
}

func (s *service) DeleteTask(task model.Task) error {
	sharevar.InfoLogger.Println("Request", task)

	err := s.repository.DeleteById(task.ID)
	if err != nil {
		sharevar.ErrorLogger.Println(err)
		return err
	}

	sharevar.InfoLogger.Println("Response")
	return nil
}
