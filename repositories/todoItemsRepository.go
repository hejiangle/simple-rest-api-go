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

func GetTodoItems() []dto.TodoItem {
	var todoItems []dto.TodoItem
	database.Find(&todoItems)

	return todoItems
}

func initToDoItems() []dto.TodoItem{
	items := append([]dto.TodoItem{},
	dto.TodoItem{ Status: false, Content: "the first test message" },
	dto.TodoItem{ Status: false, Content: "the second test message" },
	dto.TodoItem{ Status: false, Content: "the third test message" },
	)

	return items
}
