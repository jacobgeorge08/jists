package models

import "errors"

var (
	ErrNoRecord        = errors.New("models : no matching record found")
	InvalidCredentials = errors.New("models : invalid credentials")
	ErrDuplicateEmail  = errors.New("models: duplicate email")
)
