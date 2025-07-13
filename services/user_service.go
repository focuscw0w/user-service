package service

import (
	"errors"
	"log"

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

func (s *UserService) Create(req *CreateUserRequest) (*repository.User ,error) {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("username, email and password must not be empty")
	}

	user := &repository.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	return user, s.userRepo.CreateUser(user)
}

func (s *UserService) ListUsers() error {
	users, err := s.userRepo.ListUsers()
	if err != nil {
		return err
	}

	for _, u := range users {
		log.Println(u.Username)
	}

	return nil
}