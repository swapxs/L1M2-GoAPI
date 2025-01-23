package main

import (
	"log"
	"swapxs/api_proj/pkg/api"
	"swapxs/api_proj/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DBInit()
	defer database.CloseDB()

	r := gin.Default()

	if e := r.Run(":8080"); e != nil {
		log.Fatalf("Failed to start server: %v", e)
	}

	// Part 1
	r.POST("/tasks", api.Create)
	// Part 3
	r.PUT("/tasks/:id", api.Update)
	// Part 4
	r.DELETE("/tasks/:id", api.Delete)
}
