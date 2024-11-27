/*
Copyright © 2024 Kerimcan Balkan kerimcanbalkan@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command for the task management CLI application
var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A simple CLI tool to manage your daily tasks",
	Long: `Tasks is a command-line application for managing your daily to-do list.
It allows you to add, list, complete, and reset tasks efficiently. 
Organize your tasks and boost your productivity right from the terminal!`,
}

// Execute adds all child commands to the root command and handles execution.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
