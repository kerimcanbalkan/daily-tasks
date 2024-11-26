/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [task number]",
	Short: "Mark a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		taskNumber, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid task number")
			os.Exit(1)
		}
		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("error loading tasks", err)
			os.Exit(1)
		}

		if taskNumber > len(tasks) {
			fmt.Println("task not found")
			os.Exit(1)
		}

		tasks[taskNumber-1].Complete = true
		if err := SaveTasks(tasks); err != nil {
			fmt.Println("error saving task:", err)
			os.Exit(1)
		}

		fmt.Printf("Task %d marked as complete\n", taskNumber)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
