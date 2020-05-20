package repositories

import (
	"../models"
	"time"
)

func CreateNewItem(data models.TodoItemRequestModel) models.TodoItem {
	todoItem := models.TodoItem{
		Status: false,
		Content: data.Content,
		Timestamp: time.Now(),
	}

	insertStatement := `INSERT INTO todo_items VALUES ($1, $2, $3)`

	_ , err := database.Exec(insertStatement, todoItem.Status, todoItem.Content, todoItem.Timestamp)

	if err != nil {
		panic(err)
	}

	return todoItem
}
