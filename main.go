package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
}

var taskList []Task

func CreateTaskHandler(c *gin.Context) {
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

func GetTaskHandler(c *gin.Context) {
	id := c.Param("id")
	//var id string

	for _, task := range taskList {
		if strconv.Itoa(task.ID) == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func GetAllTaskHandler(c *gin.Context) {

	c.JSON(http.StatusOK, taskList)
}

func main() {

	r := gin.Default()
	r.POST("/task", CreateTaskHandler)
	r.GET("/task", GetTaskHandler)
	r.GET("/task/:id", GetTaskHandler)
	r.GET("/tasks", GetAllTaskHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
