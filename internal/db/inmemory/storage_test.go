package inmemory

import (
	"errors"
	"math"
	"testing"
	"webServerEx/internal/entity"
)

func TestStorageAdd(t *testing.T) {
	t.Run("add correct task", func(t *testing.T) {
		storage := NewStorage()
		task := &entity.Task{Title: "test"}

		err := storage.Add(task)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("add some tasks", func(t *testing.T) {
		storage := NewStorage()
		tableTests := []*entity.Task{
			{Title: "task 1"},
			{Title: "task 2"},
			{Title: "task 3"},
		}

		for _, task := range tableTests {
			err := storage.Add(task)
			if err != nil {
				t.Error(err.Error())
			}
		}
		if storage.length != 3 {
			t.Errorf("uncorrect length: %d", storage.length)
		}
		if storage.currentId != 3 {
			t.Errorf("uncorrect id %d", storage.currentId)
		}
	})

	t.Run("add nil task", func(t *testing.T) {
		storage := NewStorage()

		err := storage.Add(nil)
		if !errors.Is(err, ErrTaskIsNil) {
			t.Errorf("expected ErrTaskIsNil, got %v", err)
		}
	})

	t.Run("add maximum tasks", func(t *testing.T) {
		storage := NewStorage()
		storage.length = math.MaxUint64
		task := &entity.Task{Title: "test"}

		err := storage.Add(task)
		if !errors.Is(err, ErrTooManyTasks) {
			t.Errorf("expected ErrTooManyTasks, got %v", err)
		}
	})
}

func TestStorageGet(t *testing.T) {
	t.Run("get correct all tasks", func(t *testing.T) {
		storage := NewStorage()
		tableTests := []*entity.Task{
			{Title: "task 1"},
			{Title: "task 2"},
			{Title: "task 3"},
		}
		for _, task := range tableTests {
			storage.Add(task)
		}

		for _, task := range tableTests {
			taskFromTable, err := storage.Get(task.ID)
			if err != nil {
				t.Error(err.Error())
			}
			if taskFromTable.Title != task.Title {
				t.Errorf("uncorrect task title: %s", taskFromTable.Title)
			}
		}
	})

	t.Run("get task from wrong id", func(t *testing.T) {
		storage := NewStorage()

		_, err := storage.Get(10)
		if !errors.Is(err, ErrTaskNotFound) {
			t.Errorf("expected ErrTaskNotFound, got %v", err)
		}
	})
}

func TestStorageGetAll(t *testing.T) {
	t.Run("get all correct tasks", func(t *testing.T) {
		storage := NewStorage()
		tableTests := []*entity.Task{
			{Title: "task 1"},
			{Title: "task 2"},
			{Title: "task 3"},
		}
		for _, task := range tableTests {
			storage.Add(task)
		}

		tasks, err := storage.GetAll()
		if err != nil {
			t.Error(err.Error())
		}
		if len(tasks) != 3 {
			t.Errorf("uncorrect length from get: %d", len(tasks))
		}
	})

	t.Run("get all from empty storage", func(t *testing.T) {
		storage := NewStorage()

		tasks, err := storage.GetAll()
		if !errors.Is(err, ErrStorageEmpty) {
			t.Errorf("expected ErrStorageEmpty, got %v", err)
		}
		if len(tasks) != 0 {
			t.Errorf("uncorrect length from get: %d", len(tasks))
		}
	})
}

func TestStorageDelete(t *testing.T) {
	t.Run("delete correct", func(t *testing.T) {
		storage := NewStorage()
		task := &entity.Task{Title: "test"}
		storage.Add(task)

		err := storage.Delete(task.ID)
		if err != nil {
			t.Error(err.Error())
		}
		if storage.length != 0 {
			t.Errorf("uncorrect length: %d", storage.length)
		}
		if storage.currentId != 1 {
			t.Errorf("uncorrect currentId in storage: %d", storage.currentId)
		}
	})

	t.Run("delete wrong id task", func(t *testing.T) {
		storage := NewStorage()
		task := &entity.Task{Title: "test"}
		storage.Add(task)

		err := storage.Delete(10)
		if !errors.Is(err, ErrTaskNotFound) {
			t.Errorf("expected ErrTaskNotFound, got %v", err)
		}
		if storage.length != 1 {
			t.Errorf("uncorrect length: %d", storage.length)
		}
		if storage.currentId != 1 {
			t.Errorf("uncorrect currentId in storage: %d", storage.currentId)
		}
	})

	t.Run("delete some tasks", func(t *testing.T) {
		storage := NewStorage()
		tableTests := []*entity.Task{
			{Title: "task 1"},
			{Title: "task 2"},
			{Title: "task 3"},
		}
		for _, task := range tableTests {
			storage.Add(task)
		}

		for _, task := range tableTests {
			err := storage.Delete(task.ID)
			if err != nil {
				t.Error(err.Error())
			}
		}
		if storage.length != 0 {
			t.Errorf("uncorrect length: %d", storage.length)
		}
		if storage.currentId != 3 {
			t.Errorf("uncorrect currentId in storage: %d", storage.currentId)
		}
	})
}

func TestStorageDeleteAll(t *testing.T) {
	t.Run("delete all correct", func(t *testing.T) {
		storage := NewStorage()
		tableTests := []*entity.Task{
			{Title: "task 1"},
			{Title: "task 2"},
			{Title: "task 3"},
		}
		for _, task := range tableTests {
			storage.Add(task)
		}

		err := storage.DeleteAll()
		if err != nil {
			t.Error(err.Error())
		}
		if storage.length != 0 {
			t.Errorf("uncorrect length: %d", storage.length)
		}
		if storage.currentId != 0 {
			t.Errorf("uncorrect currentId in storage: %d", storage.currentId)
		}
	})

	t.Run("delete all from empty", func(t *testing.T) {
		storage := NewStorage()

		err := storage.DeleteAll()
		if !errors.Is(err, ErrStorageEmpty) {
			t.Errorf("expected ErrStorageEmpty, got %v", err)
		}
	})
}

func TestStorageUpdate(t *testing.T) {
	t.Run("update correct", func(t *testing.T) {
		storage := NewStorage()
		task := &entity.Task{Title: "test"}
		updatedTask := &entity.Task{Title: "update", Description: "description"}
		storage.Add(task)

		err := storage.Update(task.ID, updatedTask)
		if err != nil {
			t.Error(err.Error())
		}
		taskAfter, err := storage.Get(task.ID)
		if err != nil {
			t.Error(err.Error())
		}
		if taskAfter.Title != updatedTask.Title {
			t.Errorf("uncorrect task title: %s", updatedTask.Title)
		}
		if taskAfter.Description != updatedTask.Description {
			t.Errorf("uncorrect task description: %s", updatedTask.Description)
		}
	})

	t.Run("update wrong id", func(t *testing.T) {
		storage := NewStorage()
		updatedTask := &entity.Task{Title: "test"}

		err := storage.Update(10, updatedTask)
		if !errors.Is(err, ErrTaskNotFound) {
			t.Errorf("expected ErrTaskNotFound, got %v", err)
		}
	})

	t.Run("update task nil", func(t *testing.T) {
		storage := NewStorage()
		task := &entity.Task{Title: "test"}

		err := storage.Update(task.ID, nil)
		if !errors.Is(err, ErrTaskIsNil) {
			t.Errorf("expected ErrTaskIsNil, got %v", err)
		}
	})
}
