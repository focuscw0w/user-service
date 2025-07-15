package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/focuscw0w/microservices/services"
)

type Handler struct {
	UserService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{UserService: userService}
}

type Response struct {
	Message string `json:"message"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Rejected non-POST method: %s", r.Method)
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req service.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userDTO, err := h.UserService.CreateUser(&req)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, fmt.Sprintf("Could not create user: %v", err), http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	err = json.NewEncoder(buffer).Encode(userDTO)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	h.UserService.ListUsers()
}
