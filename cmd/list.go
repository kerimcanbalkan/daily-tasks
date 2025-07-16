/*
Copyright © 2024 Kerimcan Balkan kerimcanbalkan@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/kerimcanbalkan/daily-tasks/task"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks for that day",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []task.Task
		var err error

		if args[0] != "" && args[0] == "all" {
			tasks, err = task.GetAllTasks()
			if err != nil {
				fmt.Println("Error loading tasks", err)
				os.Exit(1)
			}
		} else {
			// Load tasks from the tasks file
			tasks, err = task.LoadTasks()
			if err != nil {
				fmt.Println("Error loading tasks", err)
				os.Exit(1)
			}
		}

		// Check if there are no tasks to display
		if len(tasks) == 0 {
			fmt.Println("No tasks found")
		}

		// Create a tabwriter for nicely formatted output
		writer := tabwriter.NewWriter(
			os.Stdout, 0, 2, 4, ' ', 0,
		)

		// Write the table header
		if _, err := writer.Write([]byte("ID\tDescription\tCreatedAt\tIsComplete\n")); err != nil {
			fmt.Println("tabwriter had an error")
			os.Exit(1)
		}

		// Loop through the tasks and format each on for display
		for _, task := range tasks {
			formattedString := fmt.Sprintf(
				"%d\t%s\t%s\t%t\n",
				task.ID,
				task.Description,
				task.CreatedAt.Local().Format("15:04"),
				task.Complete,
			)
			if _, err := writer.Write([]byte(formattedString)); err != nil {
				fmt.Println("tabwriter had an error")
				os.Exit(1)
			}
		}

		// Flush the tabwriter to ensure all data is written to output
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
