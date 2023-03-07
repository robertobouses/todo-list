package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

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

func Data(){
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
// Ejecutar el servidor Gin
if err := r.Run(":8080"); err != nil {
	log.Fatal(err)
}}

