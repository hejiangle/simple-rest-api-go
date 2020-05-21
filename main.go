package main

import (
	"./controllers"
	"./repositories"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {

	repositories.ConnectToDatabase()
	repositories.Migration()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/healthyCheck", controllers.Healthy).Methods(http.MethodGet)
	api.HandleFunc("/createItem", controllers.CreateTodoItem).Methods(http.MethodPost)
	api.HandleFunc("/todoItems", controllers.TodoItems).Methods(http.MethodGet)
	api.HandleFunc("/todoItems/{id}", controllers.GetToDoItem).Methods(http.MethodGet)

	api.HandleFunc("/", put).Methods(http.MethodPut)
	api.HandleFunc("/", delete).Methods(http.MethodDelete)
	api.HandleFunc("/", notFound)

	log.Fatal(http.ListenAndServe(":8080", r))
}
