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
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid task number")
			os.Exit(1)
		}

		err = task.CompleteTask(taskID)
		if err != nil {
			fmt.Errorf("Something Went Wrong! Task could not set complete")
		}

		// Print the succes message if no errors happened
		fmt.Printf("Task %d marked as complete\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
