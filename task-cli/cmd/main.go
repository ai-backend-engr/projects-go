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
	// Parse the flag
	flag.Parse()

	// Initializing a Tasks struct, so we can call its methods
	tasks := task.Tasks{}

	// Check if the args are passed correctly for the AddTask method,
	// call the AddTaks method, else show the cli usage and exit the program
	if os.Args[1] == "add" && len(os.Args) == 3 {
		out, err := tasks.AddTask(*add, os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Printf("# Output: %s\n", out)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
