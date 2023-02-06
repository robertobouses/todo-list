package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
}

var taskList []Task

func PostTask(c *gin.Context) {
	var task Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = len(taskList) + 1
	taskList = append(taskList, task)
	fmt.Println(taskList)

	c.JSON(http.StatusCreated, task)

}

func main() {

	r := gin.Default()
	r.POST("/task", PostTask)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
