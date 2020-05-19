package controllers

import (
	"../models"
	"encoding/json"
	"net/http"
	"time"
)

func CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/type")

	var requestModel models.TodoItemRequestModel
	err := json.NewDecoder(r.Body).Decode(&requestModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoItem := models.TodoItem{
		Status: false,
		Content: requestModel.Content,
		Timestamp: time.Now(),
	}

	response, _ := json.Marshal(todoItem)

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}