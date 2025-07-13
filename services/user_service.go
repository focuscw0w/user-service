package service

import (
	"github.com/focuscw0w/microservices/repositories"
)

// service dependency
type UserService struct {
	userRepo repository.Repository
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{userRepo: repo}
}

type CreateUserRequest struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *UserService) Create(req *CreateUserRequest) error {
	return nil
}
