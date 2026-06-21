package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ai-backend-engr/projects-go/tree/main/task-cli/internal/ui"
	task "github.com/ai-backend-engr/projects-go/tree/main/task-cli/pkg"
)

// Default filname
var filename = "./tasks.json"

func main() {
	// Parse the command flags
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	delete := flag.Int("delete", 0, "Delete a todo task")
	update := flag.Int("update", 0, "Update a todo task description")
	done := flag.Int("mark-done", 0, "Update a todo task status (done)")
	progress := flag.Int("mark-in-progress", 0, "Update a todo task status (in-progress)")

	flag.Parse()

	t := &task.Tasks{}

	// Get existing tasks from the json fiile
	if err := t.Get(filename); err != nil {
		printErr(err)
	}

	switch {
	case *add != "":
		id := t.Add(*add)
		save(*t)
		fmt.Printf("# Output: Task added successfully (ID: %d)\n", id)
	case *list:
		if len(*t) <= 0 {
			const colorReset = "\033[0m"
			const colorPrompt = "\033[32m" // Green
			const colorCmd = "\033[1;36m"  // Bold Cyan
			fmt.Printf("# Output: You currently do not have any task\n  To add a new task\n\t %s$%s %stask -add 'task name'%s\n", colorPrompt, colorReset, colorCmd, colorReset)

			os.Exit(0)
		}

		ui.RenderTable(t)
	case *done > 0:
		if err := t.Status(*done, task.StatusDone); err != nil {
			printErr(err)
		}
		save(*t)
	case *progress > 0:
		if err := t.Status(*progress, task.StatusInProgress); err != nil {
			printErr(err)
		}
		save(*t)
	case *delete > 0:
		if err := t.Delete(*delete); err != nil {
			printErr(err)
		}
		save(*t)
	case *update > 0:
		if err := t.Update(*update, flag.Arg(0)); err != nil {
			printErr(err)
		}
		save(*t)
	default:
		fmt.Fprintf(os.Stderr, "Invalid flag passed")
		flag.Usage()
		os.Exit(1)
	}
}

func printErr(err error) {
	fmt.Fprintf(os.Stderr, "%s", err)
	os.Exit(1)
}

func save(t task.Tasks) {
	if err := t.Save(filename); err != nil {
		printErr(err)
	}
}
