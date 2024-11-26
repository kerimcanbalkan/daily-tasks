/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a New task",
	Run: func(cmd *cobra.Command, args []string) {
		taskDesc := args[0]

		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks: ")
			os.Exit(1)
		}

		tasks = append(tasks, Task{Description: taskDesc, Complete: false, CreatedAt: time.Now()})
		if err := SaveTasks(tasks); err != nil {
			fmt.Println("Error saving task", err)
			os.Exit(1)
		}

		fmt.Println("Task added", taskDesc)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
