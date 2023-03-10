package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTasksNext(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

	limit := c.Query("limit")
	if limit == "" {
		limit = "10"
	}
	daysAheadStr := c.Query("daysAhead")
	daysAhead, err := strconv.Atoi(daysAheadStr)
	if err != nil {
		daysAhead = 7 // usar un valor predeterminado si daysAhead no se puede convertir a entero
	}

	now := time.Now()                      // fecha actual
	future := now.AddDate(0, 0, daysAhead) // fecha futura (por ejemplo, dentro de 7 días)

	rows, err := db.Query("SELECT * FROM tasks WHERE completed=false AND due_date BETWEEN $1 AND $2 ORDER BY due_date ASC LIMIT $3", now, future, limit)

	if err != nil {
		log.Printf("error querying for next tasks: %v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
		if err != nil {
			log.Printf("error querying for next tasks: %v \n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Printf("error querying for next tasks: %v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, tasks)
}
