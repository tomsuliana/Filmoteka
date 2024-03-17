package entity

import "errors"

var (
	ErrInvalidBirthday = errors.New("invalid birthday")
	ErrInvalidGender   = errors.New("invalid gender")
	ErrNotFound        = errors.New("actor not found")
	ErrBadContentType  = errors.New("invalid content type")
)
