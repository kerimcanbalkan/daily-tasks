package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	Description string    `json:"description"`
	Complete    bool      `json:"complete"`
	CreatedAt   time.Time `json:"createdAt"`
}

const taskFileName = ".tasks.json"

func GetTaskFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, taskFileName)
}

func LoadTasks() ([]Task, error) {
	filePath := GetTaskFilePath()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []Task{}, nil // No tasks yet
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(GetTaskFilePath(), data, 0644)
}

func ResetTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	// Get the current date at midnight
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Check if any task's creation date is before today (different day)
	reset := false
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
		if taskDate.Before(today) {
			reset = true
			break
		}
	}

	// Reset tasks if needed
	if reset {
		return SaveTasks([]Task{})
	}

	return nil
}
