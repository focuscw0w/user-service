package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/focuscw0w/microservices/user/internal/service"
)

type Handler struct {
	UserService *service.UserService
}

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the homepage!"))
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req service.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.UserService.Create(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("create user failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
