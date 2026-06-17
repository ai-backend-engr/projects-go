package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type status string

const (
	StatusDone       status = "done"
	StatusToDo       status = "todo"
	StatusInProgress status = "in-progress"
)

type task struct {
	ID          int       `json:"id"`
	Status      status    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Tasks []task

func (t *Tasks) AddTask(flag string, desc string) (string, error) {
	// Construct filepath
	filepath := filepath.Join(".", "task.json")

	// Create file if it doesn't exist
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(filepath)
		checkErr(err, fmt.Errorf("File creation failed: %s", err))

		defer file.Close()
	}

	// Read the json file
	byte, err := os.ReadFile(filepath)
	checkErr(err, fmt.Errorf("Error reading file: %s", err))

	// Convert the json to a slice
	var tasks Tasks
	err = json.Unmarshal(byte, &tasks)
	checkErr(err, fmt.Errorf("Error parsing json: %s", err))

	// Generating ID
	id := len(tasks) + 1

	// Add new task into the slice
	tasks = append(tasks, task{ID: id, Status: StatusToDo, Description: desc, CreatedAt: time.Now(), UpdatedAt: time.Now()})

	// Convert the slice back into JSON
	byte, err = json.Marshal(tasks)

	// Write into the file the new result
	os.WriteFile(filepath, byte, 0644)

	return fmt.Sprintf("Task added successfully (ID: %d)", id), nil
}

func checkErr(err error, message error) (string, error) {
	if err != nil {
		return "", message
	}

	return "", nil
}
