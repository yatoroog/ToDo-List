package scanner

import (
	"ToDoList/tasks"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	task *tasks.TaskManager
}

func NewTodo(task *tasks.TaskManager) *Todo {
	return &Todo{
		task: task,
	}
}

func (t *Todo) Start() {
	printPromt()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		inputCMD()

		if ok := scanner.Scan(); !ok {
			fmt.Println(failedToScan)
			return
		}

		inputString := scanner.Text()
		if err := t.process(inputString); err != nil {
			if err == exitCMD {
				fmt.Println("Выход из программы. До встречи!")
				return
			}
			fmt.Println(err)
		}

	}
}

func (t *Todo) process(inputString string) error {
	cmd := strings.TrimSpace(inputString)
	if cmd == "" {
		return emptyInput
	}

	switch cmd {
	case "add":
		title, text, deadline, err := AddingTask()
		if err != nil {
			return err
		}

		newTask, err := t.task.AddTask(title, text, deadline)
		if err != nil {
			return err
		}
		printCreatedTask(newTask)
	case "get":
		inputID, err := GetOne()
		if err != nil {
			return incorretInput
		}
		task, err := t.task.GetTask(inputID)
		if err != nil {
			return err
		}
		printTask(task)
	case "getAll":
		printTaskList(t.task.GetAll())
	case "del":
		inputID, err := DelOne()
		if err != nil {
			return incorretInput
		}
		if err := t.task.DeleteTask(inputID); err != nil {
			return err
		}
		printDeleted(inputID)
	case "done":
		scanner := bufio.NewScanner(os.Stdin)
		doneTask()
		if ok := scanner.Scan(); !ok {
			return incorretInput
		}
		raw := strings.TrimSpace(scanner.Text())
		inputID, err := strconv.Atoi(raw)
		if err != nil {
			return incorretInput
		}
		task, err := t.task.Done(inputID)
		if err != nil {
			return err
		}
		printDone(task)
	case "help":
		printHelp()
	case "exit", "quit":
		return exitCMD
	default:
		return IncorrectCMD

	}
	return nil
}
