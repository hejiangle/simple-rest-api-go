package main

import (
	"log"
	"net/http"
)

type server struct {

}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/type")
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message: "get called"}`))
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message: "post called}`))
	case http.MethodPut:
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case http.MethodDelete:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

}

func main() {
	s := server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
