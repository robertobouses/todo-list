package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/todo-list/domain"
)

func GetTasks(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

	rows, err := db.Query("SELECT id, title, description, due_date, completed FROM tasks ORDER BY id")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	tasks := []domain.Task{}
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}
