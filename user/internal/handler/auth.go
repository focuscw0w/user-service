package handler

import "net/http"

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the User Service API!"))
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {}

func HandleRegister(w http.ResponseWriter, r *http.Request) {}

