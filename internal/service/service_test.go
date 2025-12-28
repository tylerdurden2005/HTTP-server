package service

import (
	"errors"
	"strconv"
	"testing"
	"webServerEx/internal/entity"
)

type mockRepository struct{}

func (m mockRepository) Add(task *entity.Task) error {
	return nil
}

func (m mockRepository) Delete(id uint64) error {
	return nil
}

func (m mockRepository) DeleteAll() error {
	return nil
}

func (m mockRepository) Get(id uint64) (*entity.Task, error) {
	return nil, nil
}

func (m mockRepository) GetAll() ([]*entity.Task, error) {
	return nil, nil
}

func (m mockRepository) Update(id uint64, task *entity.Task) error {
	return nil
}

func TestServiceAddTask(t *testing.T) {
	t.Run("addTask correct tasks", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)
		tableTests := []struct {
			title       string
			description string
		}{
			{title: "title1", description: "description1"},
			{title: "title2", description: "description2"},
			{title: "d9  9o dkk 9kdkd", description: "ofkofk fj fi jik kf ikd k"},
			{title: "1234"},
			{title: "12345", description: ""},
		}

		for _, tt := range tableTests {
			err := service.AddTask(tt.title, tt.description)
			if err != nil {
				t.Error(err.Error())
			}
		}
	})

	t.Run("addTask wrong task", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)
		tableTests := []struct {
			title       string
			description string
		}{
			{title: "", description: "description1"},
			{title: "", description: "description2"},
			{title: "", description: "ofkofk fj fi jik kf ikd k"},
			{title: ""},
			{title: "", description: ""},
		}

		for _, tt := range tableTests {
			err := service.AddTask(tt.title, tt.description)
			if !errors.Is(err, ErrInvalidTitle) {
				t.Errorf("expected ErrInvalidTitle, got %v", err)
			}
		}
	})
}

func TestServiceDeleteTask(t *testing.T) {
	t.Run("deleteTask correct tasks", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)
		tableTests := []struct {
			title       string
			description string
		}{
			{title: "title1", description: "description1"},
			{title: "title2", description: "description2"},
			{title: "d9  9o dkk 9kdkd", description: "ofkofk fj fi jik kf ikd k"},
			{title: "1234"},
			{title: "12345", description: ""},
		}
		for _, tt := range tableTests {
			service.AddTask(tt.title, tt.description)
		}

		for id, _ := range tableTests {
			err := service.DeleteTask(strconv.Itoa(id))
			if err != nil {
				t.Error(err.Error())
			}
		}
	})

	t.Run("deleteTask wrong task", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)

		err := service.DeleteTask("-101")
		if !errors.Is(err, ErrInvalidID) {
			t.Errorf("expected ErrInvalidID, got %v", err)
		}
		err = service.DeleteTask("-1")
		if !errors.Is(err, ErrInvalidID) {
			t.Errorf("expected ErrInvalidID, got %v", err)
		}
	})
}

func TestServiceDeleteAllTasks(t *testing.T) {
	t.Run("deleteAllTasks correct tasks", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)
		tableTests := []struct {
			title       string
			description string
		}{
			{title: "title1", description: "description1"},
			{title: "title2", description: "description2"},
			{title: "d9  9o dkk 9kdkd", description: "ofkofk fj fi jik kf ikd k"},
			{title: "1234"},
			{title: "12345", description: ""},
		}
		for _, tt := range tableTests {
			service.AddTask(tt.title, tt.description)
		}

		err := service.DeleteAllTasks()
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestServiceGetTask(t *testing.T) {
	t.Run("getTask correct tasks", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)
		tableTests := []struct {
			title       string
			description string
		}{
			{title: "title1", description: "description1"},
			{title: "title2", description: "description2"},
			{title: "d9  9o dkk 9kdkd", description: "ofkofk fj fi jik kf ikd k"},
			{title: "1234"},
			{title: "12345", description: ""},
		}
		for _, tt := range tableTests {
			service.AddTask(tt.title, tt.description)
		}

		for id, _ := range tableTests {
			_, err := service.GetTask(strconv.Itoa(id))
			if err != nil {
				t.Error(err.Error())
			}
		}
	})

	t.Run("getTask wrong task", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)

		_, err := service.GetTask("-101")
		if !errors.Is(err, ErrInvalidID) {
			t.Errorf("expected ErrInvalidID, got %v", err)
		}
		_, err = service.GetTask("-1")
		if !errors.Is(err, ErrInvalidID) {
			t.Errorf("expected ErrInvalidID, got %v", err)
		}
	})
}

func TestServiceGetAllTasks(t *testing.T) {
	t.Run("getAllTasks correct tasks", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)
		tableTests := []struct {
			title       string
			description string
		}{
			{title: "title1", description: "description1"},
			{title: "title2", description: "description2"},
			{title: "d9  9o dkk 9kdkd", description: "ofkofk fj fi jik kf ikd k"},
			{title: "1234"},
			{title: "12345", description: ""},
		}
		for _, tt := range tableTests {
			service.AddTask(tt.title, tt.description)
		}

		_, err := service.GetAllTasks()
		if err != nil {
			t.Error(err)
		}
	})
}

func TestServiceUpdateTask(t *testing.T) {
	t.Run("updateTask correct tasks", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)
		tableTests := []struct {
			title       string
			description string
		}{
			{title: "test1", description: ""},
			{title: "test2", description: ""},
			{title: "test3", description: ""},
		}
		for _, tt := range tableTests {
			service.AddTask(tt.title, tt.description)
		}
		tableTestsUpd := []struct {
			id          string
			title       string
			description string
			finished    bool
		}{
			{id: "0", title: "update1", description: "description1", finished: true},
			{id: "1", title: "update2", description: "description2", finished: true},
			{id: "2", title: "update3"},
		}

		for _, tt := range tableTestsUpd {
			err := service.UpdateTask(tt.id, tt.title, tt.description, tt.finished)
			if err != nil {
				t.Error(err.Error())
			}
		}
	})

	t.Run("updateTask wrong task", func(t *testing.T) {
		storage := mockRepository{}
		service := NewTasksService(storage)
		tableTests := []struct {
			title       string
			description string
		}{
			{title: "test1", description: ""},
			{title: "test2", description: ""},
			{title: "test3", description: ""},
		}
		for _, tt := range tableTests {
			service.AddTask(tt.title, tt.description)
		}

		err := service.UpdateTask("-1", "update1", "update2", true)
		if !errors.Is(err, ErrInvalidID) {
			t.Errorf("expected ErrInvalidID, got %v", err)
		}
		err = service.UpdateTask("1", "", "update2", true)
		if !errors.Is(err, ErrInvalidTitle) {
			t.Errorf("expected ErrInvalidTitle, got %v", err)
		}
	})
}
