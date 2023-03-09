package handlers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func GetTasksNext(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

	limit := c.Query("limit")
	if limit == "" {
		limit = "10"
	}

	rows, err := db.Query("SELECT * FROM tasks WHERE completed=false AND due_date<=NOW() ORDER BY due_date ASC LIMIT $1", limit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, tasks)
}
