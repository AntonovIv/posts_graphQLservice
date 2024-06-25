package validation

import (
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func CreatePostValidate(post model.CreatePostReq) error {
	if post.Name == "" || post.Content == "" || post.Author == "" {
		return fmt.Errorf("forbidden empty field")
	}
	if len([]rune(post.Name)) > maxPostNameLength {
		return fmt.Errorf("invalid post name length")
	}
	if len([]rune(post.Author)) > maxAuthorLength {
		return fmt.Errorf("invalid post author length")
	}

	return nil
}
