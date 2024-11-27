/*
Copyright © 2024 Kerimcan Balkan kerimcanbalkan@gmail.com
*/
package main

import (
	"github.com/kerimcanbalkan/daily-tasks/cmd"
	"github.com/kerimcanbalkan/daily-tasks/task"
)

func main() {
	// Call ResetTasks to cleanup the outdated tasks before starting the application
	task.ResetTasks()

	// Execute the root command from the Cobra CLI framework
	// This starts the command-line application and processes user commands
	cmd.Execute()
}
