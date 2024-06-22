package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (r *repository) CreatePost(ctx context.Context, post model.CreatePostReq) (model.PostListEl, error) {
	query := `insert into Posts (name, content, author, comments_allowed)
	values ($1, $2, $3, $4)`

	_, err := r.db.DB(ctx).Exec(ctx, query, post.Name,
		post.Content, post.Author, post.CommentsAllowed)
	if err != nil {
		return model.PostListEl{}, fmt.Errorf("repository create post err: %w", err)
	}

	return model.PostListEl{}, nil
}
