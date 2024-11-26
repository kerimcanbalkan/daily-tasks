/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the current task status",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("error loading tasks")
			os.Exit(1)
		}

		incompleteCount := 0
		for _, task := range tasks {
			if !task.Complete {
				incompleteCount++
			}
		}

		if incompleteCount == 0 {
			fmt.Println("No daily tasks")
		} else {
			fmt.Printf("%d incomplete task(s)\n", incompleteCount)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
