package postgre

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (r *repository) CreatePost(ctx context.Context, post model.Post) (model.Post, error) {

	query := `insert into Posts (name, content, author, comments_allowed)
	values ($1, $2, $3, $4)`
	_, err := r.db.DB(ctx).Exec(ctx, query, post.Name,
		post.Content, post.Author, post.CommentsAllowed)

	if err != nil {
		return model.Post{}, err
	}

	return model.Post{}, err
}
