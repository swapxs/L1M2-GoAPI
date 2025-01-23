/* Created by Swapnil Bhowmik (XS/IN/0893) for Go API Task in L1: Module 2 
* This files acts as an intermediary between the main go file and the database
* functions. These parse and validate the requests from each endpoints to the
* desired function in the database.go file. */

package api

import (
	"fmt"
	"net/http"
	"strconv"
	"swapxs/api_proj/pkg/database"
	"swapxs/api_proj/pkg/format"

	"github.com/gin-gonic/gin"
)

// CRUD Operations
// 1. Create new data
func Create(c *gin.Context) {
	var newTask format.Task

	if e := c.BindJSON(&newTask); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid Request: %v", e)})
		return
	}

	t, e := database.CreateTask(newTask)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create task: %v", e)})
		return
	}

	c.JSON(http.StatusCreated, t)
}

// 2. Read specified data & ReadAll data
func Read(c *gin.Context) {
	getIdPara := c.Param("id")
	
	id, e := strconv.Atoi(getIdPara)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	t, e := database.GetTaskID(id)

	if e != nil {
		if e.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get task:\n%v", e)})
		}
		return
	}

	c.JSON(http.StatusOK, t)
}

func ReadAll(c *gin.Context) {
	t, e := database.GetAllTasks()

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get all tasks:\n%v", e)})
		return
	}

	c.JSON(http.StatusOK, t)
}

// 3. Update specified data
func Update(c *gin.Context) {
	var updateTask format.Task
	getIdPara := c.Param("id")
	
	id, e := strconv.Atoi(getIdPara)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	if e := c.BindJSON(&updateTask); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": fmt.Sprintf("Invalid request:\n%v", e)})
		return
	}

	t, e := database.UpdateTask(id, updateTask)

	if e != nil {
		if e.Error() == "Task not found" {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": fmt.Sprintf("Failed to get task:\n%v", e)})
		}
		return
	}

	c.JSON(http.StatusOK, t)
}

// 3. Delete specified data
func Delete(c *gin.Context) {
	getIdPara := c.Param("id")
	
	id, e := strconv.Atoi(getIdPara)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid task ID"})
		return
	}

	e = database.DeleteTask(id)

	if e != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Task deleted successfully"})
}
