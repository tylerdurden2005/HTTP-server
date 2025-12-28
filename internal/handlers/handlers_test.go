package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"webServerEx/internal/db/inmemory"
	"webServerEx/internal/entity"
	"webServerEx/internal/service"
)

type mockService struct {
	err error
}

func (m mockService) AddTask(title, description string) error {
	return m.err
}

func (m mockService) UpdateTask(id, title, description string, finished bool) error {
	return m.err
}

func (m mockService) GetTask(id string) (*entity.Task, error) {
	switch id {
	case "1":
		return &entity.Task{ID: 1, Title: "test"}, nil
	case "999":
		return nil, m.err
	case "2":
		return nil, m.err
	}
	return nil, nil
}

func (m mockService) GetAllTasks() ([]*entity.Task, error) {
	return nil, m.err
}

func (m mockService) DeleteTask(id string) error {
	return m.err
}

func (m mockService) DeleteAllTasks() error {
	return m.err
}

func TestHandlerGetTask(t *testing.T) {
	t.Run("handlerGet correct task", func(t *testing.T) {
		mock := mockService{}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "1")

		handler.GetTask(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status http.StatusOK, got %d", rec.Code)
		}
		var task entity.Task
		if err := json.Unmarshal(rec.Body.Bytes(), &task); err != nil {
			t.Fatalf("failed to read JSON %v", err)
		}
		if task.ID != 1 {
			t.Errorf("uncorrect id %d", task.ID)
		}
	})

	t.Run("handlerGet task fail 404", func(t *testing.T) {
		mock := mockService{err: inmemory.ErrTaskNotFound}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodGet, "/todos/999", nil)
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "999")

		handler.GetTask(rec, req)
		if rec.Code != http.StatusNotFound {
			t.Errorf("expected status http.StatusNotFound, got %d", rec.Code)
		}
	})

	t.Run("handlerGet task fail 400", func(t *testing.T) {
		mock := mockService{err: service.ErrInvalidID}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodGet, "/todos/2", nil)
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "2")

		handler.GetTask(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status http.StatusBadRequest, got %d", rec.Code)
		}
		req = httptest.NewRequest(http.MethodGet, "/todos/", nil)
		rec = httptest.NewRecorder()
		req.SetPathValue("id", "")
		handler.GetTask(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status http.StatusBadRequest, got %d", rec.Code)
		}
	})
}

func TestHandlerGetAllTasks(t *testing.T) {
	t.Run("handlerGet all tasks", func(t *testing.T) {
		mock := mockService{}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		rec := httptest.NewRecorder()

		handler.GetAllTasks(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status http.StatusOK, got %d", rec.Code)
		}
		var tasks []*entity.Task
		if err := json.Unmarshal(rec.Body.Bytes(), &tasks); err != nil {
			t.Fatalf("failed to read JSON %v", err)
		}
	})

	t.Run("handlerGet all tasks fail 400", func(t *testing.T) {
		mock := mockService{err: inmemory.ErrStorageEmpty}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		rec := httptest.NewRecorder()

		handler.GetAllTasks(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status http.StatusBadRequest, got %d", rec.Code)
		}
	})
}

func TestHandlerAddTask(t *testing.T) {
	t.Run("handlerAdd correct task", func(t *testing.T) {
		mock := mockService{err: nil}
		handler := NewHandler(mock)
		body := `{"title":"test","description":""}`
		req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.CreateTask(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status http.StatusOK, got %d", rec.Code)
		}
	})

	t.Run("handlerAdd invalid title", func(t *testing.T) {
		mock := mockService{err: service.ErrInvalidTitle}
		handler := NewHandler(mock)
		body := `{"title":""}`
		req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		handler.CreateTask(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
		}
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("handlerDelete correct task", func(t *testing.T) {
		mock := mockService{}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "1")

		handler.DeleteTask(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status http.StatusOK, got %d", rec.Code)
		}
	})

	t.Run("handlerDelete task fail 404", func(t *testing.T) {
		mock := mockService{err: inmemory.ErrTaskNotFound}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodDelete, "/todos/999", nil)
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "999")

		handler.DeleteTask(rec, req)
		if rec.Code != http.StatusNotFound {
			t.Errorf("expected status http.StatusNotFound, got %d", rec.Code)
		}
	})

	t.Run("handlerDelete task fail 400", func(t *testing.T) {
		mock := mockService{}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodDelete, "/todos/", nil)
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "")

		handler.DeleteTask(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status http.StatusBadRequest, got %d", rec.Code)
		}
	})
}

func TestHandlerDeleteAllTasks(t *testing.T) {
	t.Run("handlerDelete correct all tasks", func(t *testing.T) {
		mock := mockService{err: nil}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodDelete, "/todos", nil)
		rec := httptest.NewRecorder()

		handler.DeleteTasks(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status http.StatusOK, got %d", rec.Code)
		}
	})

	t.Run("handlerDelete all tasks fail 400", func(t *testing.T) {
		mock := mockService{err: inmemory.ErrStorageEmpty}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodDelete, "/todos", nil)
		rec := httptest.NewRecorder()

		handler.DeleteTasks(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status http.StatusBadRequest, got %d", rec.Code)
		}
	})
}

func TestHandlerUpdateTask(t *testing.T) {
	t.Run("handlerUpdate correct task", func(t *testing.T) {
		mock := mockService{}
		handler := NewHandler(mock)
		body := `{"title":"test","description":"description"}`
		req := httptest.NewRequest(http.MethodPut, "/todos/1", strings.NewReader(body))
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "1")

		handler.UpdateTask(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status http.StatusOK, got %d", rec.Code)
		}
	})

	t.Run("handlerUpdate task fail 404", func(t *testing.T) {
		mock := mockService{err: inmemory.ErrTaskNotFound}
		handler := NewHandler(mock)
		body := `{"title":"test","description":"description"}`
		req := httptest.NewRequest(http.MethodPut, "/todos/1", strings.NewReader(body))
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "1")

		handler.UpdateTask(rec, req)
		if rec.Code != http.StatusNotFound {
			t.Errorf("expected status http.StatusNotFound, got %d", rec.Code)
		}
	})

	t.Run("handlerUpdate task fail 400", func(t *testing.T) {
		mock := mockService{err: service.ErrInvalidID}
		handler := NewHandler(mock)
		req := httptest.NewRequest(http.MethodPut, "/todos/", nil)
		rec := httptest.NewRecorder()
		req.SetPathValue("id", "")

		handler.UpdateTask(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status http.StatusBadRequest, got %d", rec.Code)
		}
	})

	t.Run("handlerUpdate invalid title", func(t *testing.T) {
		mock := mockService{err: service.ErrInvalidTitle}
		handler := NewHandler(mock)
		body := `{"title":""}`
		req := httptest.NewRequest(http.MethodPut, "/todos/1", strings.NewReader(body))
		req.SetPathValue("id", "1")
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.UpdateTask(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
		}
	})
}
