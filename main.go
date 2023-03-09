package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/robertobouses/todo-list/handlers"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Completed   bool   `json:"completed"`
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
	r.POST("/tasks", handlers.PostTasks)

	// Obtener todas las tareas
	r.GET("/tasks", handlers.GetTasks)

	// Obtener una tarea por su ID
	r.GET("/tasks/:id", handlers.GetTasksId)

	// Obtener todas las tareas completadas
	r.GET("/tasks/completed", handlers.GetTasksCompleted)

	// Obtener todas las tareas no completadas
	r.GET("/tasks/pending", handlers.GetTasksPending)

	// Obtener todas las tareas con fecha expirada
	r.GET("/tasks/expired", handlers.GetTasksExpired)

	// Actualizar una tarea existente
	r.PUT("/tasks/:id", handlers.PutTasksId)

	// Obtener todas las tareas que vencen hoy
	r.GET("/tasks/today", handlers.GetTasksToday)

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
				due_date DATE,
				completed BOOLEAN NOT NULL DEFAULT false 
			);
		`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tabla creada correctamente")
}
