package controllers

import "net/http"

func Healthy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/type")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome to simple api demo", "version": "1.0"}`))
}
