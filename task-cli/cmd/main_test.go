package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "task"
	filename = "./task.json"
)

// Function that runs setup process for tests
func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	// update filename for windows OS
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	// Compiles the package
	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	// Execute the tests
	fmt.Println("Running tests")
	result := m.Run()

	// Cleaning up resources
	fmt.Println("Cleaning up...")
	os.Remove(filename)
	os.Remove(binName)

	os.Exit(result)
}

// Test your command line interface (integration test)
func TestTaskCLI(t *testing.T) {
	task := "Wake up"

	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't get the working dir: %s", err)
	}

	cmdPath := filepath.Join(dir, binName)

	// Testing task adding
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := "# Output: Task added successfully (ID: 1)\n"

		if string(out) != expected {
			t.Errorf("Expected %s, got %s", expected, string(out))
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("MarkDone", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-mark-done", "1")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("MarkInProgress", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-mark-in-progress", "1")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("UpdateTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-update", "1", "another name")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-delete", "1")

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
}
