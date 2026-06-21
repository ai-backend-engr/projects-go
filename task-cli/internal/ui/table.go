package ui

import (
	"flag"
	"fmt"

	task "github.com/ai-backend-engr/projects-go/tree/main/task-cli/pkg"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func RenderTable(t *task.Tasks) {
	// Initialize and create table title and header
	tw := table.NewWriter()
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
}
