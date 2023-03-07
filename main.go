package main

import "github.com/robertobouses/todo-list/handlers"

func main() {
	handlers.Data()
	handlers.GetTasksCompleted()
	handlers.GetTasksExpired()
	handlers.GetTasksId()
	handlers.GetTasksPending()
	handlers.GetTasks()
	handlers.PostTasks()
	handlers.PutTaskId()
}
