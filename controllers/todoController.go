package controllers

import (
	"../repositories"
	"../requestModels"
	"encoding/json"
	"net/http"
)

func CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/type")

	var requestModel requestModels.TodoItemRequestModel
	err := json.NewDecoder(r.Body).Decode(&requestModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoItem :=repositories.CreateNewItem(requestModel.Content)

	response, _ := json.Marshal(todoItem)

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func TodoItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/type")

	todoItems := repositories.GetTodoItems()

	response, _ := json.Marshal(todoItems)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}