// Package task provides functionality for managing and  reseting tasks.
// Tasks are stored as JSON in a file located in the user's home directory
package task

import (
	"log"
	"time"

	"github.com/kerimcanbalkan/daily-tasks/db"
)

// Task represents a single task with a description, completion status, and creation time.
type Task struct {
	ID          int64
	Description string
	Complete    bool
	CreatedAt   time.Time
}

func (t Task) SaveTask() (int64, error) {
	db, err := db.DBConnnect()
	if err != nil {
		return -1, err
	}
	// Use Option A (SQLite default) or B (Go UTC timestamp)
	t.CreatedAt = time.Now().UTC()
	res, err := db.Exec(
		`INSERT INTO tasks (description, complete, created_at) VALUES (?, ?, ?)`,
		t.Description, t.Complete, t.CreatedAt.Format(time.RFC3339),
	)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func CompleteTask(id int) error {
	db, err := db.DBConnnect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		"UPDATE tasks SET complete = ? WHERE id = ?",
		1, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id int) error {
	db, err := db.DBConnnect()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(
		"DELETE FROM TASKS WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func LoadTasks() ([]Task, error) {
	db, err := db.DBConnnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
  SELECT id, description, complete, created_at
  FROM tasks
  WHERE date(created_at, 'localtime') = date('now', 'localtime')
`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		var created string
		err := rows.Scan(&t.ID, &t.Description, &t.Complete, &created)
		if err != nil {
			log.Fatal(err)
		}
		t.CreatedAt, _ = time.Parse(time.RFC3339, created)
		tasks = append(tasks, t)
	}

	return tasks, nil
}
