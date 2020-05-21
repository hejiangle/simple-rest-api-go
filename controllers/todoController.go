package controllers

import (
	"../repositories"
	"../requestModels"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestModel requestModels.TodoItemRequestModel
	err := json.NewDecoder(r.Body).Decode(&requestModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoItem := repositories.CreateNewItem(requestModel.Content)

	response, _ := json.Marshal(todoItem)

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func TodoItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todoItems := repositories.GetTodoItems()

	response, _ := json.Marshal(todoItems)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetToDoItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)

	var id int
	var err error
	if value, ok := pathParams["id"]; ok {
		id, err = strconv.Atoi(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a todo item id!"}`))
			return
		}
	}

	todoItem := repositories.GetTodoItem(id)

	response, _ := json.Marshal(todoItem)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
