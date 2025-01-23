/* Created by Swapnil Bhowmik (XS/IN/0893) for Go API Task in L1: Module 2
* This files makes a struct that mimics the JSON structure for the database
* operations */

package format

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Status      string `json:"status"`
}
