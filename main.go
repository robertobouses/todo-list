package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
}

func main() {
	// Abrir la conexión con la base de datos
	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear la tabla si no existe
	createTable(db)

	// Crear el enrutador Gin
	r := gin.Default()

	// Crear una nueva tarea
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

	// Ejecutar el servidor Gin
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			due_date DATE
		);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tabla creada correctamente")
}
