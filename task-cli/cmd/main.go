package main

import (
	"flag"
	"fmt"
	"os"

	task "github.com/ai-backend-engr/projects-go/tree/main/task-cli/pkg"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	add := flag.String("add", "", "Add a new task")

	list := flag.Bool("list", false, "List all tasks")

	done := flag.Int("mark-done", 0, "Update a todo task status (done)")
	progress := flag.Int("mark-in-progress", 0, "Update a todo task status (in-progress)")

	delete := flag.Int("delete", 0, "Delete a todo task")
	update := flag.Int("update", 0, "Update a todo task description")

	// Parse the command flags
	flag.Parse()

	t := &task.Tasks{}

	// default filename
	filename := "./tasks.json"

	// Get existing tasks from the database(json fiile)
	if err := t.Get(filename); err != nil {
		printErr(err)
	}

	switch {
	case *add != "":
		id := t.Add(*add)

		if err := t.Save(filename); err != nil {
			printErr(err)
		}

		fmt.Printf("# Output: Task added successfully (ID: %d)\n", id)

	case *list:
		if len(*t) <= 0 {
			const colorReset = "\033[0m"
			const colorPrompt = "\033[32m" // Green
			const colorCmd = "\033[1;36m"  // Bold Cyan
			fmt.Printf("# Output: You currently do not have any task\n  To add a new task\n\t %s$%s %stask -add 'task name'%s\n", colorPrompt, colorReset, colorCmd, colorReset)

			os.Exit(0)
		}

		// Initialize and create table title and header
		tw := table.NewWriter()
		tw.SetTitle("All your tasks")
		tw.AppendHeader(table.Row{"ID", "DESCRIPTION", "STATUS", "CREATED AT", "UPDATED AT"})

		// Append table rows
		for _, v := range *t {
			tw.AppendRow(table.Row{v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt})
		}

		// Filter the rows, if arg is provided
		if len(flag.Args()) > 0 {
			tw.FilterBy([]table.FilterBy{{Number: 3, Operator: table.Equal, Value: flag.Arg(0)}})
		}

		// Transformer function to dynamically change the status color
		nameTransformer := text.Transformer(func(val any) string {
			switch val {
			case task.StatusToDo:
				return text.Colors{text.FgCyan}.Sprint(val)
			case task.StatusDone:
				return text.Colors{text.FgRed}.Sprint(val)
			case task.StatusInProgress:
				return text.Colors{text.FgGreen}.Sprint(val)
			}
			return fmt.Sprint(val)
		})

		tw.SetColumnConfigs([]table.ColumnConfig{
			{
				Number:      3,
				Transformer: nameTransformer,
			},
		})

		// Render the table
		fmt.Println(tw.Render())

	case *done > 0:
		if err := t.Status(*done, task.StatusDone); err != nil {
			printErr(err)
		}

		if err := t.Save(filename); err != nil {
			printErr(err)
		}

	case *progress > 0:
		if err := t.Status(*progress, task.StatusInProgress); err != nil {
			printErr(err)
		}

		if err := t.Save(filename); err != nil {
			printErr(err)
		}

	case *delete > 0:
		if err := t.Delete(*delete); err != nil {
			printErr(err)
		}

		if err := t.Save(filename); err != nil {
			printErr(err)
		}

	case *update > 0:
		t.Update(*update, flag.Arg(0))
		if err := t.Save(filename); err != nil {
			printErr(err)
		}

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
