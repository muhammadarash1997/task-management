package service

import (
	"github.com/muhammadarash1997/task-management/task/repository"
	"github.com/muhammadarash1997/task-management/task/model"
	"log"
	"runtime"
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
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request :%v", file, line, task)

	err := s.repository.CreateTask(task)
	if err != nil {
		log.Printf("Error :%v", err)
		return err
	}

	log.Printf("[%v] [%v] Response :%v", file, line, task)
	return nil
}

func (s *service) FindTasks() ([]model.Task, error) {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request", file, line)

	tasks, err := s.repository.GetAllTask()
	if err != nil {
		log.Printf("Error :%v", err)
		return []model.Task{}, err
	}

	log.Printf("[%v] [%v] Response :%v", file, line, tasks)
	return tasks, nil
}

func (s *service) FindTask(taskId uint) (model.Task, error) {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request :%v", file, line, taskId)
	
	task, err := s.repository.GetTaskById(taskId)
	if err != nil {
		log.Printf("Error :%v", err)
		return model.Task{}, err
	}

	log.Printf("[%v] [%v] Response :%v", file, line, task)
	return task, nil
}

func (s *service) UpdateTask(task model.Task) error {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request :%v", file, line, task)

	err := s.repository.UpdateTask(task)
	if err != nil {
		log.Printf("Error :%v", err)
		return err
	}

	log.Printf("[%v] [%v] Response", file, line)
	return nil
}

func (s *service) DeleteTask(task model.Task) error {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("[%v] [%v] Request :%v", file, line, task)

	err := s.repository.DeleteById(task.ID)
	if err != nil {
		log.Printf("Error :%v", err)
		return err
	}

	log.Printf("[%v] [%v] Response", file, line)
	return nil
}