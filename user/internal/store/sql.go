package store

import (
	"database/sql"

	"github.com/focuscw0w/microservices/user/internal/repository"
	
)

type SqlStorage struct {
	db *sql.DB
}

func NewSqlStorage(db *sql.DB) *SqlStorage {
	return &SqlStorage{db: db}
}

func (s *SqlStorage) CreateUser() error {
	return nil
}

func (s *SqlStorage) GetUserByID(id int) (*repository.User, error) {
	return nil, nil
}
	
func (s *SqlStorage) UpdateUser(user *repository.User) error {
	return nil
}

func (s *SqlStorage) DeleteUser(id int) error {
	return nil
}

func (s *SqlStorage) ListUsers() ([]*repository.User, error) {
	return nil, nil
}