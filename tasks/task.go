package tasks

import "time"

type Task struct {
	ID        int
	Title     string
	Text      string
	Deadline  string
	IsDone    bool
	CreatedAt time.Time
	DoneAt    *time.Time
}
