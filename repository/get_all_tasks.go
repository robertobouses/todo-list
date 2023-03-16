package repository

import "github.com/robertobouses/todo-list/domain"

func (repo TaskRepository) GetAllTasks() ([]domain.Task, error) {
	tasks := []domain.Task{}
	rows, err := repo.DB.Query("SELECT id, title, description, due_date, completed FROM tasks ORDER BY id")
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Completed); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
