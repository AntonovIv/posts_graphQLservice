package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) GetPostByID(ctx context.Context, id int) (model.Post, error) {

	query := `select id, name, author, content, comments_allowed 
	from posts where id = $1`

	var post model.Post

	err := pgxscan.Get(ctx, r.db.DB(ctx), &post, query, id)

	if pgxscan.NotFound(err) {
		return model.Post{}, models.ErrNotFound
	} else if err != nil {
		return model.Post{}, fmt.Errorf("pgxscan get: %w", err)
	}
	fmt.Println(post)
	return post, nil
}
