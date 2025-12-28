package entity

import "fmt"

type Task struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
}

func NewTask(id uint64, title string, description string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Finished:    false,
	}
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
}

func (t *Task) UpdateTitle(title string) {
	t.Title = title
}

func (t *Task) UpdateFinished() {
	t.Finished = true
}

func (t *Task) PrintTask() {
	fmt.Printf("Id:%d Title:%s Description:%s\n", t.ID, t.Title, t.Description)
}
