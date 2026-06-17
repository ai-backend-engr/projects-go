package task

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestAddTask(t *testing.T) {
	// Initialize a struct of type Tasks
	tasks := Tasks{}
	t.Run("AddTask", func(t *testing.T) {
		_, err := tasks.AddTask("add", "New task")
		if err != nil {
			log.Fatal(err)
		}

		data, err := os.ReadFile("./task.json")
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &tasks)
		if err != nil {
			log.Fatal(err)
		}

		if tasks[0].Description != "New task" && tasks[len(tasks)].ID != len(tasks)+1 {
			t.Errorf("Expected (New task, ID: 1), got %s %d", tasks[0].Description, tasks[0].ID)
		}
	})
}
