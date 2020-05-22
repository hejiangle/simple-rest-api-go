package controllers

import (
	"../repositories"
	"../requestModels"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateTodoItem(c *gin.Context) {
	w := c.Writer
	r := c.Request
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
	_, _ = w.Write(response)
}

func TodoItems(c *gin.Context) {
	w := c.Writer

	w.Header().Set("Content-Type", "application/json")

	todoItems := repositories.GetTodoItems()

	response, _ := json.Marshal(todoItems)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func GetToDoItem(c *gin.Context) {
	w := c.Writer
	r := c.Request

	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)

	id, err := resolvePathToGetId(pathParams)

	if err != nil && id == -1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoItem := repositories.GetTodoItem(id)

	response, _ := json.Marshal(todoItem)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func EditToDoItem(c *gin.Context) {
	w := c.Writer
	r := c.Request

	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)

	id, pathErr := resolvePathToGetId(pathParams)

	if pathErr != nil && id == -1 {
		http.Error(w, pathErr.Error(), http.StatusBadRequest)
		return
	}

	var requestModel requestModels.TodoItemRequestModel
	bodyErr := json.NewDecoder(r.Body).Decode(&requestModel)

	if bodyErr != nil {
		http.Error(w, bodyErr.Error(), http.StatusBadRequest)
		return
	}

	todoItem := repositories.UpdateTodoItem(id, requestModel.Content)

	response, _ := json.Marshal(todoItem)

	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write(response)
}

func DeleteToDoItem(c *gin.Context){
	w := c.Writer
	r := c.Request

	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)

	id, err := resolvePathToGetId(pathParams)

	if err != nil && id == -1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repositories.DeleteToDoItem(id)

	w.WriteHeader(http.StatusAccepted)
}

func resolvePathToGetId(params map[string]string) (int, error){
	var id int
	var err error
	if value, ok := params["id"]; ok {
		id, err = strconv.Atoi(value)
		if err != nil {
			return -1, err
		}
	}
	return id, nil
}
