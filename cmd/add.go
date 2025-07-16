/*
Copyright © 2024 Kerimcan Balkan kerimcanbalkan@gmail.com
*/
package cmd

import (
	"fmt"

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

		task := task.Task{Description: taskDesc, Complete: false}

		id, err := task.SaveTask()
		if err != nil {
			fmt.Errorf("Someting went wrong! Could not create task")
		}

		// Print a message if no errors happened
		fmt.Println("Task added with ID", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
