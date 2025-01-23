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
	"swapxs/api_proj/pkg/database"
	"swapxs/api_proj/pkg/format"
	"time"
)

// CRUD Operations
// 1. Create new data
func Create(c *gin.Context) {
	var newTask format.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("Invalid Request: %v", err)})
		return
	}

	log.Printf("Task received: %+v ", newTask)

	if err := IsValidTask(newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	t, err := database.CreateTask(newTask)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to create task: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, t)
}

// 2. Read specified data & ReadAll data
func Read(c *gin.Context) {
	getIdPara := c.Param("id")

	id, err := strconv.Atoi(getIdPara)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	t, err := database.GetTaskID(id)

	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get task: %v", err)})
		}
		return
	}

	c.JSON(http.StatusOK, t)
}

func ReadAll(c *gin.Context) {
	t, err := database.GetAllTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get all tasks: %v", err)})
		return
	}

	c.JSON(http.StatusOK, t)
}

// 3. Update specified data
func Update(c *gin.Context) {
	var updateTask format.Task
	getIdPara := c.Param("id")

	id, err := strconv.Atoi(getIdPara)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	if err := c.BindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("Invalid request: %v", err)})
		return
	}

	if err := IsValidTask(updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	t, err := database.UpdateTask(id, updateTask)

	if err != nil {
		if err.Error() == "Task not found" {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get task: %v", err)})
		}
		return
	}

	c.JSON(http.StatusOK, t)
}

// 3. Delete specified data
func Delete(c *gin.Context) {
	getIdPara := c.Param("id")

	id, err := strconv.Atoi(getIdPara)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	err = database.DeleteTask(id)

	if err != nil {
		if err.Error() == "Task not found 404" {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to delete task: %v", err)})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Task deleted successfully"})
}

func IsValidTask(t format.Task) error {
	switch {
	case strings.TrimSpace(t.Title) == "":
		return fmt.Errorf("Title cannot be empty")

	case strings.TrimSpace(t.Status) == "":
		return fmt.Errorf("Status cannot be empty")

	case strings.ToLower(t.Status) != "pending" &&
		strings.ToLower(t.Status) != "in progress" &&
		strings.ToLower(t.Status) != "completed":
		return fmt.Errorf("Invalid status, it should be 'pending', 'in progress' or 'completed'")

	case strings.TrimSpace(t.DueDate) == "":
		return fmt.Errorf("Due date cannot be empty")

	case !IsValidDate(t.DueDate):
		return fmt.Errorf("Invalid date format: Please use YYYY-MM-DD")

	default:
		return nil
	}
}

func IsValidDate(d string) bool {
	_, err := time.Parse("2006-01-02", d)

	return err == nil
}
