package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
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

func (t *Tasks) Add(desc string) int {
	var id int

	if len(*t) == 0 {
		id = 1
	} else {
		// add one to the ID of the last item
		id = (*t)[len(*t)-1].ID + 1
	}

	item := task{
		ID:          id,
		Status:      StatusToDo,
		Description: desc,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	*t = append(*t, item)

	return id
}

func (t *Tasks) Delete(id int) error {
	idx := t.getIndex(id)
	if idx == -1 {
		return fmt.Errorf("Task %d doesn't exist", id)
	}

	*t = append((*t)[:idx], (*t)[idx+1:]...)

	return nil
}

func (t *Tasks) Status(id int, status Status) error {
	idx := t.getIndex(id)
	if idx == -1 {
		return fmt.Errorf("Task %d doesn't exist", id)
	}

	(*t)[idx].Status = status

	return nil
}

func (t *Tasks) Update(id int, desc string) error {
	idx := t.getIndex(id)
	if idx == -1 {
		return fmt.Errorf("Task %d doesn't exist", id)
	}

	(*t)[idx].Description = desc
	(*t)[idx].UpdatedAt = time.Now()

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

func (t *Tasks) getIndex(id int) int {
	return slices.IndexFunc(*t, func(t task) bool {
		return t.ID == id
	})
}
