package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func get(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/type")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message: "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/type")

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message: "post called}`))
}

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

	userID := -1
	var err error
	if value, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
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
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/", post).Methods(http.MethodPost)
	api.HandleFunc("/", put).Methods(http.MethodPut)
	api.HandleFunc("/", delete).Methods(http.MethodDelete)
	api.HandleFunc("/", notFound)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
