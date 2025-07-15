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
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid task number")
			os.Exit(1)
		}

		err = task.DeleteTask(taskID)
		if err != nil {
			fmt.Errorf("Something went wrong! Could not delete the task")
		}

		// Print the succes message with deleted tasks description
		fmt.Printf("Task deleted: %d\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
