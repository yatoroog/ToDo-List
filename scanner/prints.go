package scanner

import (
	"ToDoList/tasks"
	"fmt"
	"strings"
)

func printPromt() {
	fmt.Println("")
	fmt.Println("Добро пожаловать в To Do List!")
	fmt.Println("Вы можете ввести команду >help< для просмотра действующих команд.")
	fmt.Println("")
}
func inputCMD() {
	fmt.Print("Введите команду: ")
}

func addTitle() {
	fmt.Println("Добавление задачи!")
	fmt.Print("Пожалуйста, введите название задачи: ")
}
func addText() {
	fmt.Print("Пожалуйста, введите описание задачи: ")
}
func addDeadline() {
	fmt.Print("Пожалуйста, укажите дедлайн в формате ДД.ММ.ГГ или укажите '-', если дедлайн не нужен!")
}

func getTask() {
	fmt.Print("Пожалуйста, укажите ID задачи, которую хотите посмотреть: ")
}
func delTask() {
	fmt.Print("Пожалуйста, укажите ID задачи, которую хотите удалить: ")
}
func doneTask() {
	fmt.Print("Пожалуйста, укажите ID задачи, которую хотите отметить как 'выполненную': ")
}

func printHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println(" add    - добавить задачу")
	fmt.Println(" get    - показать задачу по ID")
	fmt.Println(" getAll - показать все задачи")
	fmt.Println(" del    - удалить задачу по ID")
	fmt.Println(" done   - отметить задачу как выполненную")
	fmt.Println(" help   - показать список команд")
	fmt.Println(" exit   - выйти из программы")
}

func printCreatedTask(task tasks.Task) {
	fmt.Printf("Задача создана. ID: %d, название: %s\n", task.ID, task.Title)
}

func printTaskList(taskList []tasks.Task) {
	if len(taskList) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}

	fmt.Printf("Всего задач: %d\n", len(taskList))
	fmt.Println(strings.Repeat("-", 30))
	for _, task := range taskList {
		printTask(task)
	}
}

func printTask(task tasks.Task) {
	status := "Не выполнена"
	doneAt := "-"
	if task.IsDone {
		status = "Выполнена"
		if task.DoneAt != nil {
			doneAt = task.DoneAt.Format("02.01.2006 15:04")
		}
	}

	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("Название: %s\n", task.Title)
	fmt.Printf("Описание: %s\n", task.Text)
	fmt.Printf("Дедлайн: %s\n", task.Deadline)
	fmt.Printf("Статус: %s\n", status)
	fmt.Printf("Создано: %s\n", task.CreatedAt.Format("02.01.2006 15:04"))
	fmt.Printf("Завершено: %s\n", doneAt)
	fmt.Println(strings.Repeat("-", 30))
}

func printDeleted(id int) {
	fmt.Printf("Задача с ID %d удалена.\n", id)
}

func printDone(task tasks.Task) {
	fmt.Printf("Задача с ID %d отмечена как выполненная.\n", task.ID)
}
