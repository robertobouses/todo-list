// Obtener todas las tareas
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks() {
	r.GET("/tasks", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, description, due_date, completed FROM tasks ORDER BY id")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer rows.Close()

		tasks := []Task{}
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			tasks = append(tasks, task)
		}

		c.JSON(http.StatusOK, tasks)
	})
}
