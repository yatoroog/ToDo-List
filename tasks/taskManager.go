package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type TaskManager struct {
	tasks    map[int]Task
	nextID   int
	filePath string
}

func NewTaskManager() (*TaskManager, error) {
	tm := &TaskManager{
		tasks:    make(map[int]Task),
		nextID:   1,
		filePath: defaultFilePath(),
	}

	if err := tm.loadFromFile(); err != nil {
		return nil, fmt.Errorf("%w: %v", failedToLoad, err)
	}

	return tm, nil
}

func (tm *TaskManager) AddTask(title, text, deadline string) (Task, error) {
	if strings.TrimSpace(title) == "" {
		return Task{}, emptyTitle
	}
	task := Task{
		ID:        tm.nextID,
		Title:     title,
		Text:      text,
		Deadline:  deadline,
		IsDone:    false,
		CreatedAt: time.Now(),
		DoneAt:    nil,
	}

	tm.tasks[tm.nextID] = task
	tm.nextID++

	if err := tm.saveToFile(); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (tm *TaskManager) GetTask(id int) (Task, error) {
	task, ok := tm.tasks[id]
	if !ok {
		return Task{}, notFound
	}
	return task, nil
}

func (tm *TaskManager) GetAll() []Task {
	list := make([]Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		list = append(list, task)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list
}

func (tm *TaskManager) DeleteTask(id int) error {
	if _, exist := tm.tasks[id]; !exist {
		return notFound
	}
	delete(tm.tasks, id)
	return tm.saveToFile()
}

func (tm *TaskManager) Done(id int) (Task, error) {
	task, ok := tm.tasks[id]
	if !ok {
		return Task{}, notFound
	}
	if task.IsDone {
		return Task{}, alreadyDone
	} else {
		task.IsDone = true
		now := time.Now()
		task.DoneAt = &now
		tm.tasks[id] = task
		if err := tm.saveToFile(); err != nil {
			return Task{}, err
		}
		return task, nil
	}

}

func defaultFilePath() string {
	return "tasks.json"
}

func (tm *TaskManager) loadFromFile() error {
	file, err := os.Open(tm.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	var data struct {
		Tasks  []Task `json:"tasks"`
		NextID int    `json:"next_id"`
	}

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return err
	}

	for _, task := range data.Tasks {
		tm.tasks[task.ID] = task
	}

	if data.NextID > 0 {
		tm.nextID = data.NextID
	} else {
		maxID := 0
		for id := range tm.tasks {
			if id > maxID {
				maxID = id
			}
		}
		tm.nextID = maxID + 1
		if tm.nextID == 0 {
			tm.nextID = 1
		}
	}

	return nil
}

func (tm *TaskManager) saveToFile() error {
	taskList := tm.GetAll()

	payload := struct {
		Tasks  []Task `json:"tasks"`
		NextID int    `json:"next_id"`
	}{
		Tasks:  taskList,
		NextID: tm.nextID,
	}

	content, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Errorf("%w: %v", failedToSave, err)
	}

	tmpPath := tm.filePath + ".tmp"
	if err := os.WriteFile(tmpPath, content, 0o644); err != nil {
		return fmt.Errorf("%w: %v", failedToSave, err)
	}

	if err := os.Rename(tmpPath, tm.filePath); err != nil {
		return fmt.Errorf("%w: %v", failedToSave, err)
	}

	return nil
}
