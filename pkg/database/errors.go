package database

import "errors"

var (
	ErrNotFound      = errors.New("entity not found")
	ErrAlreadyExists = errors.New("entity is already exists")
)
