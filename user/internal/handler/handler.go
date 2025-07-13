package handler

import (
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

	var req service.CreateUserRequest

	h.UserService.Create(&req)
}
