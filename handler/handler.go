package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/focuscw0w/microservices/internal/security"
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

func (h *Handler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Rejected non-POST method: %s", r.Method)
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req service.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userDTO, err := h.UserService.SignUp(&req)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, fmt.Sprintf("Could not create user: %v", err), http.StatusInternalServerError)
		return
	}

	token, err := security.CreateToken(userDTO.Username)
	if err != nil {
		log.Printf("Failed to create session token: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

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

func (h *Handler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Rejected non-POST method: %s", r.Method)
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req service.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userDTO, err := h.UserService.SignIn(&req)
	if err != nil {
		log.Printf("Failed to sign in user: %v", err)
		http.Error(w, fmt.Sprintf("Could not sign in user: %v", err), http.StatusUnauthorized)
		return
	}

	token, err := security.CreateToken(userDTO.Username)
	if err != nil {
		log.Printf("Failed to create session token: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	buffer := new(bytes.Buffer)
	err = json.NewEncoder(buffer).Encode(userDTO)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Rejected non-GET method: %s", r.Method)
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	usersDTO, err := h.UserService.ListUsers()
	if err != nil {
		log.Printf("Failed to list users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	err = json.NewEncoder(buffer).Encode(usersDTO)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}
