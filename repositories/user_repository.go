package repository

import "database/sql"

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

type Repository interface {
	CreateUser(user *User) error
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

func (s *SqlStorage) CreateUser(user *User) error {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	res, err := s.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)

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
	query := `SELECT * FROM users`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
