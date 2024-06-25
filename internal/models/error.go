package models

import (
	"errors"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrBadPostId = errors.New("bad request")
)
