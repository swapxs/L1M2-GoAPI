/* Created by Swapnil Bhowmik (XS/IN/0893) for Go API Task in L1: Module 2
* Main file for the entire project, this file performs the following:
* 1. Initializes the database
* 2. Creates a Router using gin
* 3. Performs all the required tasks outlined in the document
* 4. Runs the server and checks for errors.
* 5. Closes the Database after it is finished executing.*/

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/swapxs/GoAPI/pkg/api"
	"github.com/swapxs/GoAPI/pkg/db"
)

func main() {
	db.Init()
	defer db.Close()

	r := gin.Default()

	r.POST("/tasks", api.Create)
	r.GET("/tasks/:id", api.Read)
	r.PUT("/tasks/:id", api.Update)
	r.DELETE("/tasks/:id", api.Delete)
	r.GET("/tasks", api.ReadAll)

	if e := r.Run(":8080"); e != nil {
		log.Fatalf("Failed to start server: %v", e)
	}
}
