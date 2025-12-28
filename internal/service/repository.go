package service

import "webServerEx/internal/entity"

type Repository interface {
	Add(task *entity.Task) error
	Delete(id uint64) error
	DeleteAll() error
	Get(id uint64) (*entity.Task, error)
	GetAll() ([]*entity.Task, error)
	Update(id uint64, task *entity.Task) error
}

type tasksRepository struct {
	storage Repository
}

func NewRepository(storage Repository) Repository {
	return &tasksRepository{storage: storage}
}

func (r *tasksRepository) Add(task *entity.Task) error {
	return r.storage.Add(task)
}

func (r *tasksRepository) Delete(id uint64) error {
	return r.storage.Delete(id)
}
func (r *tasksRepository) DeleteAll() error {
	return r.storage.DeleteAll()
}
func (r *tasksRepository) Get(id uint64) (*entity.Task, error) {
	return r.storage.Get(id)
}
func (r *tasksRepository) GetAll() ([]*entity.Task, error) {
	return r.storage.GetAll()
}
func (r *tasksRepository) Update(id uint64, task *entity.Task) error {
	return r.storage.Update(id, task)
}
