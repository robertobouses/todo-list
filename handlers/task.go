package handlers

import "github.com/robertobouses/todo-list/repository"

type taskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) taskHandler {
	return taskHandler{
		repo: repo,
	}
}
