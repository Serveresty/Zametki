package customerrors

import "errors"

var ErrNoTokenFound = errors.New("no token found")
var ErrInvalidToken = errors.New("invalid token")
var ErrAlreadyRegistered = errors.New("already registered")
var ErrAlreadyAuthorized = errors.New("already authorized")
