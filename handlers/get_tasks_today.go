package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTasksToday(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

	// Obtener la fecha actual
	today := time.Now().Format("2006-01-02")

	// Realizar la consulta a la base de datos
	rows, err := db.Query("SELECT * FROM tasks WHERE due_date=$1", today)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Crear un slice de tareas
	var tasks []Task

	// Iterar sobre los resultados y agregarlos al slice
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	// Devolver el slice de tareas como respuesta
	c.JSON(http.StatusOK, tasks)
}
