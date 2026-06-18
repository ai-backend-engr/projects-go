package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Status string

const (
	StatusDone       Status = "done"
	StatusToDo       Status = "todo"
	StatusInProgress Status = "in-progress"
)

type task struct {
	ID          int       `json:"id"`
	Status      Status    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Tasks []task

func (t *Tasks) Add(desc string) {
	item := task{
		ID:          len(*t) + 1,
		Status:      StatusToDo,
		Description: desc,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	*t = append(*t, item)
}

func (t *Tasks) Delete(id int) error {
	if id <= 0 || id >= len(*t) {
		return fmt.Errorf("Task %d doesn't exist", id)
	}

	// Spread it because append
	// accepts a slice and objects/types
	*t = append((*t)[:id], (*t)[id:]...)

	return nil
}

func (t *Tasks) Status(id int, status Status) error {
	if id <= 0 || id >= len(*t) {
		return fmt.Errorf("Task %d doesn't exist", id)
	}

	(*t)[id].Status = status

	return nil
}

func (t *Tasks) Save(filepath string) error {
	js, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, js, 0644)
}

func (t *Tasks) Get(filepath string) error {
	byte, err := os.ReadFile(filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	return json.Unmarshal(byte, t)
}
