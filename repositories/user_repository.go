package repository

import "database/sql"

type Repository interface {
	GetUserByID(id int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetAllUsers() ([]*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

type SqlStorage struct {
	db *sql.DB
}

func NewSqlStorage(db *sql.DB) *SqlStorage {
	return &SqlStorage{db: db}
}

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func (s *SqlStorage) GetUserByID(id int) (*User, error) {
	query := `SELECT * FROM users WHERE id = ?`

	var u User
	row := s.db.QueryRow(query, id)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *SqlStorage) GetUserByUsername(username string) (*User, error) {
	query := `SELECT * FROM users WHERE username = ?`

	var u User
	row := s.db.QueryRow(query, username)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *SqlStorage) GetAllUsers() ([]*User, error) {
	query := `SELECT * FROM users`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User
	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

func (s *SqlStorage) CreateUser(user *User) (*User, error) {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	res, err := s.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)

	return user, nil
}

func (s *SqlStorage) UpdateUser(user *User) error {
	return nil
}

func (s *SqlStorage) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
