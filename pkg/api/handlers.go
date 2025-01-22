package api

import (
	"fmt"
	"net/http"
	"strconv"
	"swapxs/api_proj/pkg/database"
	"swapxs/api_proj/pkg/models"

	"github.com/gin-gonic/gin"
)

// CRUD
func CreateTask(c *gin.Context) {
	var newTask models.Task

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

func Update(c *gin.Context) {
	var updateTask models.Task
	getIdPara := c.Param("id")
	
	id, e := strconv.Atoi(getIdPara)

	if e != nil {
		// error
		return
	}

	t, e := database.UpdateTask(id, updateTask)

	if e != nil {
		return
	}

	c.JSON(http.StatusOK, t)
}

func Delete(c *gin.Context) {
	getIdPara := c.Param("id")
	
	id, e := strconv.Atoi(getIdPara)

	if e != nil {
		// error
		return
	}

	e = database.DeleteTask(id)

	if e != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
