// Obtener una tarea por su ID
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasksId() {
	r.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		var task Task
		err := db.QueryRow("SELECT id, title, description, due_date, completed FROM tasks WHERE id=$1", id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, task)
	})
}
