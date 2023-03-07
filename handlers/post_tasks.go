// Crear una nueva tarea
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTasks() {
	r.POST("/tasks", func(c *gin.Context) {
		var task Task
		if err := c.BindJSON(&task); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Insertar la tarea en la base de datos
		stmt, err := db.Prepare("INSERT INTO tasks(title, description, due_date) VALUES ($1, $2, $3) RETURNING id")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer stmt.Close()

		var id int
		if err := stmt.QueryRow(task.Title, task.Description, task.DueDate).Scan(&id); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Actualizar el ID de la tarea y devolverla como respuesta
		task.ID = id
		c.JSON(http.StatusOK, task)
	})
}
