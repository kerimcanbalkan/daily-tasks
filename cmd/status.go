/*
Copyright © 2024 Kerimcan Balkan kerimcanbalkan@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kerimcanbalkan/daily-tasks/task"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the current task status",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the tasks from the tasks file
		tasks, err := task.LoadTasks()
		if err != nil {
			fmt.Println("error loading tasks")
			os.Exit(1)
		}

		// Count the incomplete tasks
		incompleteCount := 0
		for _, task := range tasks {
			if !task.Complete {
				incompleteCount++
			}
		}

		// Print how many incomplete tasks are there
		if incompleteCount == 0 {
			fmt.Println("No incomplete tasks")
		} else {
			fmt.Printf("%d incomplete task(s)\n", incompleteCount)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
