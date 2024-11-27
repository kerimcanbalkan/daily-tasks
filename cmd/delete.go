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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
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

		// Check if the provided task number is within the range of available tasks
		if taskNumber > len(tasks) {
			fmt.Println("task not found")
			os.Exit(1)
		}

		// Get the description of the task provided by the argument
		taskDesc := tasks[taskNumber-1].Description

		// Delete the task from the list of tasks
		tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)

		// Save remaining tasks to the file
		if err := task.SaveTasks(tasks); err != nil {
			fmt.Println("error saving tasks: ", err)
			os.Exit(1)
		}

		// Print the succes message with deleted tasks description
		fmt.Printf("Task deleted: %s\n", taskDesc)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
