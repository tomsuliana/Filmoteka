package entity

import "errors"

var (
	ErrInvalidBirthday    = errors.New("invalid birthday")
	ErrInvalidGender      = errors.New("invalid gender")
	ErrNotFound           = errors.New("actor not found")
	ErrBadContentType     = errors.New("invalid content type")
	ErrInvalidReleaseDate = errors.New("invalid release date")
	ErrInvalidRating      = errors.New("invalid rating")
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidDescription = errors.New("invalid description")
)
