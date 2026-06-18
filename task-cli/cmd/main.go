package main

import (
	"flag"
	"fmt"
	"os"

	task "github.com/ai-backend-engr/projects-go/tree/main/task-cli/pkg"
)

func main() {
	// Get the arguments
	add := flag.String("add", "", "Add a new task")
	list := flag.String("list", "", "List all todo tasks")
	// Parse the flag
	flag.Parse()

	// Initializing a Tasks struct, so we can call its methods
	tasks := task.Tasks{}

	// Check if the args are passed correctly for the specific method
	if os.Args[1] == *add && len(os.Args) == 3 {
		// call the AddTaks method
		out, err := tasks.AddTask(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Printf("# Output: %s\n", out)
	} else if os.Args[1] == *list && len(os.Args) == 2 {
		out, err := tasks.ListTasks()
		if err != nil {
			panic(err)
		}
		fmt.Printf("All your todos\n %v", out)
	} else {
		flag.Usage()
		os.Exit(1)
	}

}
