package validation

import (
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func CreatePostValidate(post model.CreatePostReq) error {
	if post.Name == "" || post.Content == "" || post.Author == "" {
		return fmt.Errorf("forbidden empty field")
	}
	if len([]rune(post.Name)) > 255 {
		return fmt.Errorf("invalid post name length")
	}
	if len([]rune(post.Author)) > 64 {
		return fmt.Errorf("invalid post author length")
	}

	return nil
}
