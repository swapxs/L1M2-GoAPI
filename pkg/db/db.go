/* Created by Swapnil Bhowmik (XS/IN/0893) for Go API Task in L1: Module 2
* This file/module has the following function:
* 1. Handle the starting and closing of the sqlite database
* 2. perform CRUD operation by getting input from the /pkg/api/functions.go
*    file */

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/swapxs/GoAPI/pkg/format"
)

var db *sql.DB

// Start the Database
func Init() {
	var e error
	db, e = sql.Open("sqlite3", "tasks.db")

	if e != nil {
		log.Fatalf("Failed to open database: %v", e)
	}

	createTab := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			description TEXT,
			dueDate TEXT,
			status TEXT
		)
	`
	_, e = db.Exec(createTab)

	if e != nil {
		log.Fatalf("\nFailed to create table: %v\n", e)
	}
}

// Close the Database
func Close() {
	if db != nil {
		db.Close()
	}
}

func GetAllTasks() ([]format.Task, error) {
	var allTasks []format.Task

	r, e := db.Query("SELECT id, title, description, dueDate, status FROM tasks")

	if e != nil {
		return nil, fmt.Errorf("Failed to get all tasks.\nERROR: %v", e)
	}

	defer r.Close()

	for r.Next() {
		var t format.Task

		e := r.Scan(&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Status)

		if e != nil {
			return nil, fmt.Errorf("Failed to scan tasks.\nERROR: %v", e)
		}

		allTasks = append(allTasks, t)
	}
	return allTasks, nil
}

func GetTaskID(id int) (format.Task, error) {
	var t format.Task

	e := db.QueryRow("SELECT id, title, description, dueDate, status FROM tasks WHERE id = ?", id).Scan(
		&t.ID, &t.Title, &t.Description, &t.DueDate, &t.Status)

	if e == sql.ErrNoRows {
		return format.Task{}, fmt.Errorf("Task not found")
	} else if e != nil {
		return format.Task{}, fmt.Errorf("Failed to get task:\nERROR: %v", e)
	}

	return t, nil
}

// Function that creates tasks
func CreateTask(t format.Task) (format.Task, error) {
	res, e := db.Exec("INSERT INTO tasks (title, description, dueDate, status) VALUES (?, ?, ?, ?)",
		t.Title, t.Description, t.DueDate, t.Status)

	if e != nil {
		return format.Task{}, fmt.Errorf("\nFailed to create new task.\nERROR: %v\n", e)
	}

	id, e := res.LastInsertId()

	if e != nil {
		return format.Task{}, fmt.Errorf("\nFailed to get task id.\nERROR: %v\n", e)
	}

	t.ID = int(id)
	return t, nil
}

func UpdateTask(id int, t format.Task) (format.Task, error) {
	_, e := db.Exec("UPDATE tasks SET title = ?, description = ?, dueDate = ? WHERE id = ?",
		t.Title, t.Description, t.DueDate, id)

	if e != nil {
		return format.Task{}, fmt.Errorf("failed to update task: %v", e)
	}

	updatedTask, e := GetTaskID(id)

	if e != nil {
		return format.Task{}, fmt.Errorf("failed to get updated task: %v", e)
	}

	return updatedTask, nil
}

// Function that deletes tasks
func DeleteTask(id int) error {
	res, e := db.Exec("DELETE FROM tasks WHERE id = ?", id)

	if e != nil {
		return fmt.Errorf("\nFailed to delete task.\nERROR: %v", e)
	}

	found, _ := res.RowsAffected()

	if found == 0 {
		return fmt.Errorf("Task not found 404\nERROR: %v", e)
	}

	return nil
}
