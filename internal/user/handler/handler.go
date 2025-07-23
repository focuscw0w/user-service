package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	email "github.com/focuscw0w/microservices/internal/email/service"
	"github.com/focuscw0w/microservices/internal/user/security"
	user "github.com/focuscw0w/microservices/internal/user/service"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	UserService  *user.Service
	EmailService *email.Service
}

func NewHandler(userService *user.Service, emailService *email.Service) *Handler {
	return &Handler{UserService: userService, EmailService: emailService}
}

func (h *Handler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Rejected non-POST method: %s", r.Method)
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req user.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userDTO, err := h.UserService.SignUp(&req)
	if err != nil {
		log.Printf("Failed to sign up: %v", err)
		http.Error(w, "Failed to sign up", http.StatusInternalServerError)
		return
	}

	token, err := security.CreateToken(userDTO.Username)
	if err != nil {
		log.Printf("Failed to create session token: %v", err)
		http.Error(w, "Internal server errors", http.StatusInternalServerError)
		return
	}

	c := http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &c)

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

	var req user.SignInRequest
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
		http.Error(w, "Internal server errors", http.StatusInternalServerError)
		return
	}

	c := http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &c)

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

func (h *Handler) HandleSignOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Rejected non-POST method: %s", r.Method)
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	c := &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	}

	http.SetCookie(w, c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message":"Signed out successfully"}`))
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

	usersDTO, err := h.UserService.GetUsers()
	if err != nil {
		log.Printf("Failed to list users: %v", err)
		http.Error(w, "Internal server errors", http.StatusInternalServerError)
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

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Rejected non-GET method: %s", r.Method)
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Failed to parse id: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userDTO, err := h.UserService.GetUser(id)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
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
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		log.Printf("Rejected non-DELETE method: %s", r.Method)
		http.Error(w, "Only DELETE method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Failed to parse id: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.UserService.DeleteUser(id)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	c := &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	}

	http.SetCookie(w, c)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"message":"User deleted"}`))
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func (h *Handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		log.Printf("Rejected non-PUT method: %s", r.Method)
		http.Error(w, "Only PUT method is allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Failed to parse id: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req user.UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userDTO, err := h.UserService.UpdateUser(id, req)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}
