package controllers

import (
	_ "../dto"
	"../repositories"
	"../applicationModels"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @tags TodoList
// @Summary Create a new to do item
// @Produce json
// @Param content body applicationModels.TodoItemRequestModel true "Create item"
// @Success 200 {object} dto.TodoItem
// @Router /todoItems/ [post]
func CreateTodoItem(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var requestModel applicationModels.TodoItemRequestModel
	err := c.ShouldBindJSON(&requestModel)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	todoItem := repositories.CreateNewItem(requestModel.Content)

	c.JSON(http.StatusCreated, todoItem)
}

// @tags TodoList
// @Summary Get all to do items
// @Produce json
// @Success 200 {array} dto.TodoItem
// @Router /todoItems [get]
func TodoItems(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	todoItems := repositories.GetTodoItems()

	c.JSON(http.StatusOK, todoItems)
}

// @tags TodoList
// @Summary Get the to do item by ID
// @Produce json
// @Param id path int true "To do item ID" Format(int64)
// @Success 200 {object} dto.TodoItem
// @Router /todoItems/{id} [get]
func GetToDoItem(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	pathParams := c.Param("id")

	id, err := resolvePathToGetId(pathParams)

	if err != nil && id == -1 {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	todoItem := repositories.GetTodoItem(id)

	c.JSON(http.StatusOK, todoItem)
}

// @tags TodoList
// @Summary Edit the to do item by ID
// @Produce json
// @Param id path int true "To do item ID" Format(int64)
// @Param content body applicationModels.TodoItemRequestModel true "Edit item by id"
// @Success 202 {object} dto.TodoItem
// @Router /todoItems/{id} [put]
func EditToDoItem(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	pathParams := c.Param("id")

	id, pathErr := resolvePathToGetId(pathParams)

	if pathErr != nil && id == -1 {
		http.Error(c.Writer, pathErr.Error(), http.StatusBadRequest)
		return
	}

	var requestModel applicationModels.TodoItemRequestModel
	bodyErr := c.ShouldBindJSON(&requestModel)

	if bodyErr != nil {
		http.Error(c.Writer, bodyErr.Error(), http.StatusBadRequest)
		return
	}

	todoItem := repositories.UpdateTodoItem(id, requestModel.Content)

	c.JSON(http.StatusAccepted, todoItem)
}

// @tags TodoList
// @Summary Delete the to do item by ID
// @Produce json
// @Param id path int true "To do item ID" Format(int64)
// @Success 204 {object} dto.TodoItem
// @Router /todoItems/{id} [delete]
func DeleteToDoItem(c *gin.Context){
	c.Header("Content-Type", "application/json")
	pathParams := c.Param("id")

	id, err := resolvePathToGetId(pathParams)

	if err != nil && id == -1 {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	repositories.DeleteToDoItem(id)

	c.JSON(http.StatusNoContent, gin.H{})
}

func resolvePathToGetId(params string) (int, error){
	var id int
	var err error
	id, err = strconv.Atoi(params)
	if err != nil {
		return -1, err
	}
	return id, nil
}
