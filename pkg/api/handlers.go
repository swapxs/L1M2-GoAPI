package api

import (
	"fmt"
	"net/http"
	"swapxs/api_proj/pkg/database"
	"swapxs/api_proj/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if e := c.BindJSON(&newTask); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid Request: %v", e)})
		return
	}

	t, e := database.InsertTask(newTask)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create task: %v", e)})
		return
	}

	c.JSON(http.StatusCreated, t)
}
