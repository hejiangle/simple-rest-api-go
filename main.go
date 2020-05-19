package main

import (
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func put(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/type")

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/type")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/type")

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func params(w http.ResponseWriter, r *http.Request){
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/type")

	var userID int
	var err error
	if value, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	var commentID int
	if value, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID"": %d, "CommentID": %d, "location": %s }`, userID, commentID, location)))
}

func main() {

	connectToDatabase()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/healthyCheck", controllers.Healthy).Methods(http.MethodGet)
	api.HandleFunc("/createItem", controllers.CreateTodoItem).Methods(http.MethodPost)
	api.HandleFunc("/", put).Methods(http.MethodPut)
	api.HandleFunc("/", delete).Methods(http.MethodDelete)
	api.HandleFunc("/", notFound)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
