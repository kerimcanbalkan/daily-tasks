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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
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

		taskDesc := tasks[taskNumber-1].Description
		tasks = append(tasks[:taskNumber-1], tasks[taskNumber:]...)

		if err := SaveTasks(tasks); err != nil {
			fmt.Println("error saving tasks: ", err)
			os.Exit(1)
		}

		fmt.Printf("Task deleted: %s\n", taskDesc)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
