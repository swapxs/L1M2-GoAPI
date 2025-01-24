/* Created by Swapnil Bhowmik (XS/IN/0893) for Go API Task in L1: Module 2
* This files acts as an intermediary between the main go file and the database
* functions. These parse and validate the requests from each endpoints to the
* desired function in the database.go file. */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/swapxs/GoAPI/pkg/db"
	"github.com/swapxs/GoAPI/pkg/format"
	"time"
)

// Data Validation
func IsValidDate(d string) bool {
	_, e := time.Parse("2006-01-02", d)

	return e == nil
}

func IsValidTask(task format.Task) error {
	switch {
	case strings.TrimSpace(task.Title) == "":
		return fmt.Errorf("Title cannot be empty")

	case strings.TrimSpace(task.Status) == "":
		return fmt.Errorf("Status cannot be empty")

	case strings.ToLower(task.Status) != "pending" &&
		strings.ToLower(task.Status) != "in progress" &&
		strings.ToLower(task.Status) != "completed":
		return fmt.Errorf("Invalid status, it should be 'pending', 'in progress' or 'completed'")

	case strings.TrimSpace(task.DueDate) == "":
		return fmt.Errorf("Due date cannot be empty")

	case !IsValidDate(task.DueDate):
		return fmt.Errorf("Invalid date format: Please use YYYY-MM-DD")

	default:
		return nil
	}
}

// CRUD Operations
// 1. Create new data
func Create(c *gin.Context) {
	var newTask format.Task

	if e := c.BindJSON(&newTask); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("Invalid Request: %v", e)})
		return
	}

	if e := IsValidTask(newTask); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": e.Error()})
		return
	}

	log.Printf("Task received: %+v ", newTask)

	task, e := db.CreateTask(newTask)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to create task: %v", e)})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// 2. Read specified data & ReadAll data
func Read(c *gin.Context) {
	getId := c.Param("id")

	id, e := strconv.Atoi(getId)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	task, e := db.GetTaskID(id)

	if e != nil {
		if e.Error() == "Task not found" {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get task: %v", e)})
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

func ReadAll(c *gin.Context) {
	task, e := db.GetAllTasks()

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get all tasks: %v", e)})
		return
	}

	c.JSON(http.StatusOK, task)
}

// 3. Update specified data
func Update(c *gin.Context) {
	var updateTask format.Task
	getId := c.Param("id")

	id, e := strconv.Atoi(getId)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	if e := c.BindJSON(&updateTask); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("Invalid request: %v", e)})
		return
	}

	if e := IsValidTask(updateTask); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": e.Error()})
		return
	}

	task, e := db.UpdateTask(id, updateTask)

	if e != nil {
		if e.Error() == "Task not found" {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get task: %v", e)})
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

// 3. Delete specified data
func Delete(c *gin.Context) {
	getId := c.Param("id")

	id, e := strconv.Atoi(getId)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	e = db.DeleteTask(id)

	if e != nil {
		if e.Error() == "Task not found 404" {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to delete task: %v", e)})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Task deleted successfully"})
}
