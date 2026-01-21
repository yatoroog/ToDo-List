package scanner

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func AddingTask() (title, text, deadline string, err error) {
	//Добавление тайтла задачи
	scanner := bufio.NewScanner(os.Stdin)
	addTitle()
	if ok := scanner.Scan(); !ok {
		return "", "", "", emptyInput
	}
	inputTitle := strings.TrimSpace(scanner.Text())
	if inputTitle == "" {
		return "", "", "", emptyInput
	}

	//Добавление описания задачи
	scanner = bufio.NewScanner(os.Stdin)
	addText()
	if ok := scanner.Scan(); !ok {
		return "", "", "", emptyInput
	}
	inputText := strings.TrimSpace(scanner.Text())

	//Добавление дедлайна
	scanner = bufio.NewScanner(os.Stdin)
	addDeadline()
	if ok := scanner.Scan(); !ok {
		return "", "", "", emptyInput
	}
	inputDeadline := strings.TrimSpace(scanner.Text())

	title = inputTitle
	text = inputText
	deadline = inputDeadline

	return title, text, deadline, nil
}

func GetOne() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	getTask()
	if ok := scanner.Scan(); !ok {
		return 0, emptyInput
	}
	raw := strings.TrimSpace(scanner.Text())
	inputID, err := strconv.Atoi(raw)
	if err != nil {
		return 0, incorretInput
	}
	if inputID <= 0 {
		return 0, incorretInput
	}
	return inputID, nil
}
func DelOne() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	delTask()
	if ok := scanner.Scan(); !ok {
		return 0, incorretInput
	}
	raw := strings.TrimSpace(scanner.Text())
	inputID, err := strconv.Atoi(raw)
	if err != nil {
		return 0, incorretInput
	}
	if inputID <= 0 {
		return 0, incorretInput
	}
	return inputID, nil
}
