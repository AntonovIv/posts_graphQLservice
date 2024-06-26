package validation

import (
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func CreateCommentValidate(comment model.CreateCommentReq) error {
	if comment.Author == "" || comment.Content == "" {
		return fmt.Errorf("forbidden empty field")
	}

	if len([]rune(comment.Content)) > MaxCommentLength {
		return fmt.Errorf("invalid post content length")
	}
	if len([]rune(comment.Author)) > MaxAuthorLength {
		return fmt.Errorf("invalid post author length")
	}

	return nil
}
