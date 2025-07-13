package repository

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
