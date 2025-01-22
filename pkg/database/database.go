package database

import (
	"database/sql"
	"fmt"
	"log"
	"swapxs/api_proj/pkg/models"
)

var db *sql.DB

// DB Initialize/Close
func DBInit() {
	var e error
	db, e := sql.Open("sqlite3", "tasks.db")

	if e != nil {
		log.Fatalf("failed to open database: %v", e)
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
	_, e = db.Exec(tC)

	if e != nil {
		log.Fatalf("\nFailed to create table: %v\n", e)
	}
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func CreateTask(t models.Task) (models.Task, error) {
	res, e := db.Exec("INSERT INTO tasks (title, description, dueDate, status) VALUES (?, ?, ?, ?)",
		t.Title, t.Description, t.DueDate, t.Status)

	if e != nil {
		return models.Task{}, fmt.Errorf("\nFailed to create new task.\nERROR: %v\n", e)
	}

	id, e := res.LastInsertId()

	if e != nil {
		return models.Task{}, fmt.Errorf("\nFailed to get task id.\nERROR: %v\n", e)
	}

	t.ID = int(id)
	return t, nil
}

func DeleteTask(id int) error {
	_, e := db.Exec("DELETE FROM tasks WHERE id = ?", id)

	if e != nil {
		return fmt.Errorf("\nFailed to delete task.\nERROR: %v", e)
	}

	return nil
}
