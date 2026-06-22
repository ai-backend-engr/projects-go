package task_test

import (
	"os"
	"testing"
	"time"

	task "github.com/ai-backend-engr/projects-go/tree/main/task-cli/pkg"
)

func TestAdd(t *testing.T) {
	tsk := task.Tasks{}

	desc := "New Task"

	tsk.Add(desc)

	if tsk[0].Description != desc {
		t.Errorf("Expected %s, got %s", desc, tsk[0].Description)
	}
}

func TestDelete(t *testing.T) {
	tsk := task.Tasks{
		{
			ID:          1,
			Status:      task.StatusToDo,
			Description: "Task 1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Time{},
		},
		{
			ID:          2,
			Status:      task.StatusToDo,
			Description: "Task 2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Time{},
		},
		{
			ID:          3,
			Status:      task.StatusToDo,
			Description: "Task 3",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Time{},
		},
	}

	err := tsk.Delete(2)
	if err != nil {
		t.Fatalf("%s", err)
	}

	if len(tsk) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(tsk))
	}
}

func TestStatus(t *testing.T) {
	tsk := task.Tasks{}

	desc := "New Task"

	tsk.Add(desc)

	if tsk[0].Status != task.StatusToDo {
		t.Errorf("Expected %s, new task cannot have a status of %s", task.StatusToDo, tsk[0].Status)
	}

	tsk.Status(1, task.StatusInProgress)

	if tsk[0].Status != task.StatusInProgress {
		t.Errorf("Expected %s, got %s", task.StatusInProgress, tsk[0].Status)
	}

	tsk.Status(1, task.StatusDone)

	if tsk[0].Status != task.StatusDone {
		t.Errorf("Expected %s, got %s", task.StatusDone, tsk[0].Status)
	}
}

func TestUpdate(t *testing.T) {
	desc := "Before update"
	newDesc := "After update"

	tsk := task.Tasks{}

	tsk.Add(desc)

	if tsk[0].Description != desc {
		t.Errorf("Expected %s but got %s", desc, tsk[0].Description)
	}

	tsk.Update(1, newDesc)

	if tsk[0].Description != newDesc {
		t.Errorf("Expected %s but got %s", newDesc, tsk[0].Description)
	}
}

func TestSaveGet(t *testing.T) {
	tsk := task.Tasks{}
	tsk2 := task.Tasks{}

	// Create a temp file
	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Couldn't create file: %s", err)
	}

	defer os.Remove(file.Name())

	tsk.Add("Task 1")

	err = tsk.Save(file.Name())
	if err != nil {
		t.Fatalf("File not saved: %s", err)
	}

	err = tsk2.Get(file.Name())
	if err != nil {
		t.Fatalf("Error fetching file: %s", err)
	}

	if tsk[0].Description != tsk2[0].Description {
		t.Errorf("Task %q should match task %q", tsk[0].Description, tsk2[0].Description)
	}
}
