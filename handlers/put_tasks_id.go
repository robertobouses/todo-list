package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/todo-list/domain"
)

func PutTasksId(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

	id := c.Param("id")
	fmt.Println("EL VALOR DEL ID!!!!!!!!!!!!!!!", id)
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println("COMPLETED ANTES DE LA ACTUALIZACIÓN:", task.Completed)

	// Actualizar la tarea en la base de datos
	stmt, err := db.Prepare("UPDATE tasks SET title=$1, description=$2, due_date=$3, completed=$4 WHERE id=$5")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer stmt.Close()
	fmt.Printf("UPDATE tasks SET title=%s, description=%s, due_date=%s, completed=%t WHERE id=%s\n", task.Title, task.Description, task.DueDate, task.Completed, id)
	_, err = stmt.Exec(task.Title, task.Description, task.DueDate, task.Completed, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("COMPLETED DESPUÉS DE LA ACTUALIZACIÓN:", task.Completed)

	c.Status(http.StatusOK)
}
