/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
		}

		writer := tabwriter.NewWriter(
			os.Stdout, 0, 2, 4, ' ', 0,
		)

		if _, err := writer.Write([]byte("ID\tDescription\tCreatedAt\tIsComplete\n")); err != nil {
			fmt.Println("tabwriter had an error")
			os.Exit(1)
		}

		for i, task := range tasks {
			formattedString := fmt.Sprintf(
				"%d\t%s\t%s\t%t\n",
				i+1,
				task.Description,
				task.CreatedAt.Format("15:04"),
				task.Complete,
			)
			if _, err := writer.Write([]byte(formattedString)); err != nil {
				fmt.Println("tabwriter had an error")
				os.Exit(1)
			}
		}

		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
