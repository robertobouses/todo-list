package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/robertobouses/todo-list/handlers"
	"github.com/robertobouses/todo-list/repository"

	"github.com/robertobouses/todo-list/handlers/users"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Completed   bool   `json:"completed"`
}

func Login(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user struct {
		ID       int
		Password string
	}
	err = db.QueryRow("SELECT id, password FROM users WHERE email = $1", credentials.Email).Scan(&user.ID, &user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
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

func main() {
	// Crear un nuevo repositorio
	repo, err := repository.NewTaskRepository("postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// Crear la tabla si no existe
	createTable(repo.DB)

	// crear los handlers
	taskHandler := handlers.NewTaskHandler(repo)

	// Crear el enrutador Gin
	r := gin.Default()

	// Crear una nueva tarea
	r.POST("/tasks", handlers.PostTasks)

	// Obtener todas las tareas
	r.GET("/tasks", handlers.GetTasks)

	// Obtener una tarea por su ID
	r.GET("/tasks/:id", handlers.GetTasksId)

	// Obtener todas las tareas completadas
	r.GET("/tasks/completed", taskHandler.GetTasksCompleted)

	// Obtener todas las tareas no completadas
	r.GET("/tasks/pending", handlers.GetTasksPending)

	// Obtener todas las tareas con fecha expirada
	r.GET("/tasks/expired", handlers.GetTasksExpired)

	// Actualizar una tarea existente
	r.PUT("/tasks/:id", handlers.PutTasksId)

	// Obtener todas las tareas que vencen hoy
	r.GET("/tasks/today", handlers.GetTasksToday)

	// Obtener las próximas tareas no completadas
	r.GET("/tasks/next", handlers.GetTasksNext)

	// Ejecutar el servidor Gin
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	// Crear un nuevo usuario
	r.POST("/users", users.PostUsers)

	/*// Obtener todos los usuarios
		r.GET("/users", users.GetUsers)

		// Obtener un usuario por su ID
		r.GET("/users/:id", users.GetUsersId)

		// Actualizar un usuario existente
		r.PUT("/users/:id", users.PutUsersId)

		// Eliminar un usuario
		r.DELETE("/users/:id", users.DeleteUsersId)

		// Iniciar sesión (autenticación)
		r.POST("/login", users.Login)

	}*/
}
