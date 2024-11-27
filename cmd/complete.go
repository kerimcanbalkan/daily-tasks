/*
Copyright © 2024 Kerimcan Balkan kerimcanbalkan@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/kerimcanbalkan/daily-tasks/task"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [task number]",
	Short: "Mark a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure a valid task number is provided as an argument
		taskNumber, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid task number")
			os.Exit(1)
		}

		// Load the list of tasks from the tasks file
		tasks, err := task.LoadTasks()
		if err != nil {
			fmt.Println("error loading tasks", err)
			os.Exit(1)
		}

		// Check if the provided number is within the range of available tasks
		if taskNumber > len(tasks) {
			fmt.Println("task not found")
			os.Exit(1)
		}

		// Mark the specified task as complete
		tasks[taskNumber-1].Complete = true

		// Save the updated tasks
		if err := task.SaveTasks(tasks); err != nil {
			fmt.Println("error saving task:", err)
			os.Exit(1)
		}

		// Print the succes message if no errors happened
		fmt.Printf("Task %d marked as complete\n", taskNumber)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
