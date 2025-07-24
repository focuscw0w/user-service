package errors

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrEmptyCredentials = errors.New("empty user credentials")
	ErrUpdateUserFailed = errors.New("update user failed")
	ErrDeleteUserFailed = errors.New("delete user failed")
)
