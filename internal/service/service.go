package service

import (
	"errors"
	"log"
	"strconv"
	"webServerEx/internal/entity"
)

var (
	ErrInvalidTitle = errors.New("invalid title")
	ErrInvalidID    = errors.New("invalid input id")
)

type Service interface {
	AddTask(title, description string) error
	UpdateTask(id, title, description string, finished bool) error
	GetTask(id string) (*entity.Task, error)
	GetAllTasks() ([]*entity.Task, error)
	DeleteTask(id string) error
	DeleteAllTasks() error
}

type tasksService struct {
	repository Repository
}

func NewTasksService(repository Repository) Service {
	return &tasksService{repository: repository}
}

func (r *tasksService) AddTask(title, description string) error {
	if title == "" {
		log.Printf("---Service: failed to add task: %v", ErrInvalidTitle)
		return ErrInvalidTitle
	}
	task := &entity.Task{
		Title:       title,
		Description: description,
	}
	if err := r.repository.Add(task); err != nil {
		log.Printf("---Service: failed to add task to repository: %v", err)
		return err
	}
	log.Println("---Service: task added successfully")
	return nil
}

func (r *tasksService) UpdateTask(id, title, description string, finished bool) error {
	correctID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Printf("---Service: failed to update task: %v", ErrInvalidID)
		return ErrInvalidID
	}
	if title == "" {
		log.Printf("---Service: failed to update task: %v", ErrInvalidTitle)
		return ErrInvalidTitle
	}
	task := &entity.Task{
		Title:       title,
		Description: description,
		Finished:    finished,
	}
	if err = r.repository.Update(correctID, task); err != nil {
		log.Printf("---Service: failed to update task to repository: %v", err)
		return err
	}
	log.Println("---Service: task updated successfully")
	return nil
}
func (r *tasksService) GetTask(id string) (*entity.Task, error) {
	correctID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Printf("---Service: failed to get task: %v", ErrInvalidID)
		return nil, ErrInvalidID
	}
	task, err := r.repository.Get(correctID)
	if err != nil {
		log.Printf("---Service: failed to get task from repository: %v", err)
		return nil, err
	}
	log.Println("---Service: task got successfully")
	return task, nil
}
func (r *tasksService) GetAllTasks() ([]*entity.Task, error) {
	tasks, err := r.repository.GetAll()
	if err != nil {
		log.Printf("---Service: failed to get all tasks from repository: %v", err)
		return nil, err
	}
	log.Println("---Service: tasks got successfully")
	return tasks, nil
}

func (r *tasksService) DeleteTask(id string) error {
	correctID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Printf("---Service: failed to delete task: %v", ErrInvalidID)
		return ErrInvalidID
	}
	if err = r.repository.Delete(correctID); err != nil {
		log.Printf("---Service: failed to delete task from repository: %v", err)
		return err
	}
	log.Println("---Service: task deleted successfully")
	return nil
}
func (r *tasksService) DeleteAllTasks() error {
	if err := r.repository.DeleteAll(); err != nil {
		log.Printf("---Service: failed to delete all tasks from repository: %v", err)
		return err
	}
	log.Println("---Service: tasks deleted successfully")
	return nil
}
