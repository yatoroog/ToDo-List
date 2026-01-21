package main

import (
	"ToDoList/scanner"
	"ToDoList/tasks"
	"fmt"
)

func main() {
	task, err := tasks.NewTaskManager()
	if err != nil {
		fmt.Println(err)
		return
	}
	todo := scanner.NewTodo(task)
	todo.Start()
}
