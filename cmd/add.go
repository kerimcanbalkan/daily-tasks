/*
Copyright © 2024 Kerimcan Balkan kerimcanbalkan@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/kerimcanbalkan/daily-tasks/task"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a New task",
	Run: func(cmd *cobra.Command, args []string) {
		// Get task description given by user
		taskDesc := args[0]

		// Load tasks from the tasks file
		tasks, err := task.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks: ")
			os.Exit(1)
		}

		// Add new task to the tasks
		tasks = append(
			tasks,
			task.Task{Description: taskDesc, Complete: false, CreatedAt: time.Now()},
		)
		if err := task.SaveTasks(tasks); err != nil {
			fmt.Println("Error saving task", err)
			os.Exit(1)
		}

		// Print a message if no errors happened
		fmt.Println("Task added", taskDesc)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
