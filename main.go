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

	// Obtener todas las tareas
	r.GET("/tasks", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, description, due_date, completed FROM tasks ORDER BY id")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer rows.Close()

		tasks := []Task{}
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			tasks = append(tasks, task)
		}

		c.JSON(http.StatusOK, tasks)
	})

	// Obtener una tarea por su ID
	r.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		var task Task
		err := db.QueryRow("SELECT id, title, description, due_date, completed FROM tasks WHERE id=$1", id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, task)
	})

	// Obtener todas las tareas completadas
	r.GET("/tasks/completed", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, description, due_date, completed FROM tasks WHERE completed=true ORDER BY id")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer rows.Close()

		tasks := []Task{}
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			tasks = append(tasks, task)
		}

		c.JSON(http.StatusOK, tasks)
	})

	// Actualizar una tarea existente
	r.PUT("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("EL VALOR DEL ID!!!!!!!!!!!!!!!", id)
		var task Task
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
			due_date DATE,
			completed BOOLEAN NOT NULL DEFAULT false 
		);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tabla creada correctamente")
}
