package models

import "time"

type TodoItem struct {
	Status bool
	Content string
	Timestamp time.Time
}

type TodoItemRequestModel struct {
	Content string
}
