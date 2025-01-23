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

	// Part 1
	r.POST("/tasks", api.Create)

	// Part 2
	r.GET("/tasks/:id", api.Read)

	// Part 3
	r.PUT("/tasks/:id", api.Update)

	// Part 4
	r.DELETE("/tasks/:id", api.Delete)

	// Part 5
	r.GET("/tasks", api.ReadAll)

	if e := r.Run(":8080"); e != nil {
		log.Fatalf("Failed to start server: %v", e)
	}
}
