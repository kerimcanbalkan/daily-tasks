// Package task provides functionality for managing and  reseting tasks.
// Tasks are stored as JSON in a file located in the user's home directory
package task

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Task represents a single task with a description, completion status, and creation time.
type Task struct {
	Description string    `json:"description"`
	Complete    bool      `json:"complete"`
	CreatedAt   time.Time `json:"createdAt"`
}

const taskFileName = ".tasks.json"

// GetTaskFilePath returns the full path to the file where tasks are stored.
// It locates the file int the user's home directory.
func GetTaskFilePath() string {
	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	// Combine home directory path with the task file name
	return filepath.Join(home, taskFileName)
}

// LoadTasks reads the task file and returns a list of tasks.
// If the task file does not exist, it returns an empty task list
func LoadTasks() ([]Task, error) {
	// Get  the path to the task file
	filePath := GetTaskFilePath()

	// Check if the task file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []Task{}, nil // No tasks yet
	}

	// Read the file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into a slice of tasks
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// SaveTasks writes the provided list of tasks to the task file.
// Tasks are saved in a human-readable indented JSON format.
func SaveTasks(tasks []Task) error {
	// Marshal the tasks inot an indented JSON string
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	// Write the JSON data to the task file
	return os.WriteFile(GetTaskFilePath(), data, 0644)
}

// ResetTasks removes tasks that are outdated, retaining only tasks created today.
// A task is considered outdated if its creation date is before the current day.
func ResetTasks() error {
	// Load the existing tasks from the task file
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	// Get the current date at midnight
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Filter tasks to retain only the ones created on or after today
	var updatedTasks []Task
	for _, task := range tasks {
		taskDate := time.Date(
			task.CreatedAt.Year(),
			task.CreatedAt.Month(),
			task.CreatedAt.Day(),
			0,
			0,
			0,
			0,
			task.CreatedAt.Location(),
		)
		if !taskDate.Before(today) {
			updatedTasks = append(updatedTasks, task)
		}
	}

	// Save the filtered tasks
	return SaveTasks(updatedTasks)
}
