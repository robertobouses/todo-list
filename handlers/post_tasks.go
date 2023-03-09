package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Completed   bool   `json:"completed"`
}

func PostTasks(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

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
}
