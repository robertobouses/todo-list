package repository

import (
	"database/sql"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(connString string) (TaskRepository, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return TaskRepository{}, err
	}
	return TaskRepository{
		DB: db,
	}, nil
}
