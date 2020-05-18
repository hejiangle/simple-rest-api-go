package main

import (
	"log"
	"net/http"
)

type server struct {

}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/type")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message: "hello, world!"}`))
}

func main() {
	s := server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
