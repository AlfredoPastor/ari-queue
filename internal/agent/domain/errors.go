package domain

import "errors"

var ErrInvalidCredentials error = errors.New("invalid credentials")
var ErrNotFound error = errors.New("not found")
