package repository

import "database/sql"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Repository interface {
	CreateUser() error
	GetUserByID(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
	ListUsers() ([]*User, error)
}

type SqlStorage struct {
	db *sql.DB
}

func NewSqlStorage(db *sql.DB) *SqlStorage {
	return &SqlStorage{db: db}
}

func (s *SqlStorage) CreateUser() error {
	return nil
}

func (s *SqlStorage) GetUserByID(id int) (*User, error) {
	return nil, nil
}

func (s *SqlStorage) UpdateUser(user *User) error {
	return nil
}

func (s *SqlStorage) DeleteUser(id int) error {
	return nil
}

func (s *SqlStorage) ListUsers() ([]*User, error) {
	return nil, nil
}
