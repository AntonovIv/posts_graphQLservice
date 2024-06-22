package validation

import (
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func CreateCommentValidate(comment model.CreateCommentReq) error {
	if comment.Author == "" || comment.Content == "" || comment.Post == "" {
		return fmt.Errorf("forbidden empty field")
	}
	if len([]rune(comment.Post)) <= 0 {
		return fmt.Errorf("invalid post name length")
	}
	if len([]rune(comment.Content)) > 2000 {
		return fmt.Errorf("invalid post author length")
	}
	if len([]rune(comment.Author)) > 64 {
		return fmt.Errorf("invalid post author length")
	}

	return nil
}
