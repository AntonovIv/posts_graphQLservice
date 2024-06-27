package models

import (
	"errors"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrBadPostId = errors.New("bad request")

	//resolver errors
	ErrInternalServerResolver = &gqlerror.Error{
		Message: "internal server error",
	}
	ErrNotFoundResolver = &gqlerror.Error{
		Message: "posts not found",
	}
	ErrCommentsNotAllowedResolver = &gqlerror.Error{
		Message: "bad request: commenting not alowed",
	}
)
