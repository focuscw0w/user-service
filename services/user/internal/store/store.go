package store

type User struct {
	ID	  int
	Name  string
	Email string
}

type Storage interface {
	CreateUser(user *User) error
	GetUserByID(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
	ListUsers() ([]*User, error)
}