package inmemory

import (
	"errors"
	"math"
	"sync"
	"webServerEx/internal/entity"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrStorageEmpty = errors.New("tasks storage is empty")
	ErrTaskIsNil    = errors.New("task is nil")
	ErrTooManyTasks = errors.New("task is too many")
)

type TasksStorage struct {
	length    uint64
	currentId uint64
	data      map[uint64]*entity.Task
	mu        sync.RWMutex
}

func NewStorage() *TasksStorage {
	return &TasksStorage{
		length:    0,
		currentId: 0,
		data:      make(map[uint64]*entity.Task),
	}
}

func (ts *TasksStorage) Add(task *entity.Task) error {
	if task == nil {
		return ErrTaskIsNil
	}
	if ts.length == math.MaxUint64 {
		return ErrTooManyTasks
	}
	if ts.currentId == math.MaxUint64 {
		return ErrTooManyTasks
	}
	ts.mu.Lock()
	ts.data[ts.currentId] = task
	task.ID = ts.currentId
	ts.currentId++
	ts.length++
	ts.mu.Unlock()
	return nil
}

func (ts *TasksStorage) Delete(id uint64) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	if _, ok := ts.data[id]; ok {
		delete(ts.data, id)
		ts.length--
		return nil
	}
	return ErrTaskNotFound
}

func (ts *TasksStorage) DeleteAll() error {
	if ts.length == 0 {
		return ErrStorageEmpty
	}
	ts.mu.Lock()
	ts.data = make(map[uint64]*entity.Task)
	ts.length = 0
	ts.currentId = 0
	ts.mu.Unlock()
	return nil
}

func (ts *TasksStorage) Update(id uint64, task *entity.Task) error {
	if task == nil {
		return ErrTaskIsNil
	}
	ts.mu.Lock()
	defer ts.mu.Unlock()
	if _, ok := ts.data[id]; ok {
		ts.data[id] = task
		task.ID = id
		return nil
	}
	return ErrTaskNotFound
}

func (ts *TasksStorage) Get(id uint64) (*entity.Task, error) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	if v, ok := ts.data[id]; ok {
		return v, nil
	}
	return nil, ErrTaskNotFound
}

func (ts *TasksStorage) GetAll() ([]*entity.Task, error) {
	if ts.length == 0 {
		return nil, ErrStorageEmpty
	}
	data := make([]*entity.Task, 0, ts.length)
	ts.mu.RLock()
	for _, task := range ts.data {
		data = append(data, task)
	}
	ts.mu.RUnlock()
	return data, nil
}

func (ts *TasksStorage) PrintAll() {
	for _, task := range ts.data {
		task.PrintTask()
	}
}
