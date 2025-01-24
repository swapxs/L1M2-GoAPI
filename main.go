/* Created by Swapnil Bhowmik (XS/IN/0893) for Go API Task in L1: Module 2
* Main file for the entire project, this file performs the following:
* 1. Initializes the database
* 2. Creates a Router r using gin
* 3. Performs all the required tasks outlined in the document
* 4. Runs the server and checks for errors.
* 5. Closes the Database after it is finished executing.*/

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"swapxs/GoAPI/pkg/api"
	"swapxs/GoAPI/pkg/database"
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

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
