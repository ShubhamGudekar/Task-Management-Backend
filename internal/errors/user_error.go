package user_errors

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrEmailAlreadyRegistered = errors.New("email already registered")
