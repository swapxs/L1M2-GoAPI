/* Created by Swapnil Bhowmik (XS/IN/0893) for Go API Task in L1: Module 2
* This file/module has the following function:
* 1. Handle the starting and closing of the sqlite database
* 2. perform CRUD operation by getting input from the /pkg/api/functions.go
*    file */

package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"swapxs/GoAPI/pkg/format"
)

var db *sql.DB

// Start the Database
func DBInit() {
	var err error
	db, err = sql.Open("sqlite3", "tasks.db")

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	tC := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			description TEXT,
			dueDate TEXT,
			status TEXT
		)
	`
	_, err = db.Exec(tC)

	if err != nil {
		log.Fatalf("\nFailed to create table: %v\n", err)
	}
}

// Close the Database
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func GetAllTasks() ([]format.Task, error) {
	var allTasks []format.Task

	r, err := db.Query("SELECT id, title, description, dueDate, status FROM tasks")

	if err != nil {
		return nil, fmt.Errorf("Failed to get all tasks.\nERROR: %v", err)
	}

	defer r.Close()

	for r.Next() {
		var t format.Task

		err := r.Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Status)

		if err != nil {
			return nil, fmt.Errorf("Failed to scan tasks.\nERROR: %v", err)
		}

		allTasks = append(allTasks, t)
	}
	return allTasks, nil
}

func GetTaskID(id int) (format.Task, error) {
	var t format.Task

	err := db.QueryRow("SELECT id, title, description, dueDate, status FROM tasks WHERE id = ?", id).Scan(
		&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Status)

	if err == sql.ErrNoRows {
		return format.Task{}, fmt.Errorf("Task not found")
	} else if err != nil {
		return format.Task{}, fmt.Errorf("Failed to get task:\nERROR: %v", err)
	}

	return t, nil
}

// Function that creates tasks
func CreateTask(t format.Task) (format.Task, error) {
	res, err := db.Exec("INSERT INTO tasks (title, description, dueDate, status) VALUES (?, ?, ?, ?)",
		t.Title, t.Description, t.DueDate, t.Status)

	if err != nil {
		return format.Task{}, fmt.Errorf("\nFailed to create new task.\nERROR: %v\n", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return format.Task{}, fmt.Errorf("\nFailed to get task id.\nERROR: %v\n", err)
	}

	t.ID = int(id)
	return t, nil
}

func UpdateTask(id int, t format.Task) (format.Task, error) {
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, dueDate = ? WHERE id = ?",
		t.Title, t.Description, t.DueDate, id)

	if err != nil {
		return format.Task{}, fmt.Errorf("failed to update task: %v", err)
	}

	updatedTask, err := GetTaskID(id)

	if err != nil {
		return format.Task{}, fmt.Errorf("failed to get updated task: %v", err)
	}

	return updatedTask, nil
}

// Function that deletes tasks
func DeleteTask(id int) error {
	res, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		return fmt.Errorf("\nFailed to delete task.\nERROR: %v", err)
	}

	found, _ := res.RowsAffected()

	if found == 0 {
		return fmt.Errorf("Task not found 404\nERROR: %v", err)
	}

	return nil
}
