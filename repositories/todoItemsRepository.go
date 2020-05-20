package repositories

import "../dto"

func CreateNewItem(content string) dto.TodoItem {
	todoItem := dto.TodoItem{
		Status: false,
		Content: content,
	}

	if database.NewRecord(todoItem) {
		database.Create(&todoItem)
	}

	return todoItem
}

func Migration() {
	database.DropTableIfExists(&dto.TodoItem{})
	database.AutoMigrate(&dto.TodoItem{})
}
