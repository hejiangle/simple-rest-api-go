package repositories

import (
	"../dto"
)

func CreateNewItem(content string) dto.TodoItem {
	todoItem := dto.TodoItem{
		Status:  false,
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

func GetTodoItem(id int) dto.TodoItem {
	var item dto.TodoItem
	database.Where("id = ?", id).Find(&item)

	return item
}

func UpdateTodoItem(id int, content string, status bool) dto.TodoItem {
	var item dto.TodoItem
	database.Find(&item, "id = ?", id)

	if content != "" && content != item.Content && status != item.Status {
		database.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{"content": content, "status": status})
		return item
	}

	if content != "" && content != item.Content {
		database.Model(&item).Where("id = ?", id).Update("content", content)
		return item
	}

	if status != item.Status {
		database.Model(&item).Where("id = ?", id).Update("status", status)
	}

	return item
}

func DeleteToDoItem(id int) {
	database.Where("id = ?", id).Delete(&dto.TodoItem{})
}

func initToDoItems() []dto.TodoItem {
	items := append([]dto.TodoItem{},
		dto.TodoItem{Status: false, Content: "the first test message"},
		dto.TodoItem{Status: false, Content: "the second test message"},
		dto.TodoItem{Status: false, Content: "the third test message"},
	)

	return items
}
