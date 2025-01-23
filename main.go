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

	r.POST("/tasks", api.Create)
	r.PUT("/tasks/:id", api.Update)
	r.DELETE("/tasks/:id", api.Delete)
}
