package internal

import (
	"github.com/focuscw0w/microservices/user/internal/repository"
)

// service dependency
type UserService struct {
	userRepo repository.Repository
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{userRepo: repo}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *UserService) Create(req *CreateUserRequest) error {
	return nil
}
