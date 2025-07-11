package internal

import (
	"github.com/focuscw0w/microservices/services/user/internal/store"
)

// service dependency
type UserService struct {
	userRepo store.Storage
}

func NewUserService(repo store.Storage) *UserService {
	return &UserService{userRepo: repo}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *UserService) Create(req *CreateUserRequest) error {}
