package service

import (
	"errors"
	"log"

	"github.com/focuscw0w/microservices/internal/security"
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
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (s *UserService) CreateUser(req *CreateUserRequest) (*UserDTO, error) {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("username, email and password must not be empty")
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &repository.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	userDTO := &UserDTO{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}

	return userDTO, nil
}

func (s *UserService) ListUsers() error {
	users, err := s.userRepo.ListUsers()
	if err != nil {
		return err
	}

	for _, u := range users {
		log.Println(u.Password)
	}

	// transform to dto

	return nil
}
