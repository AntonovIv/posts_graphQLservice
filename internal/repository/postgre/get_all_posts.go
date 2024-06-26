package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) GetAllPosts(ctx context.Context) ([]model.PostListEl, error) {
	query := `select id, name, author, content 
	from posts`
	var postAll []model.PostListEl

	err := pgxscan.Select(ctx, r.db.DB(ctx), &postAll, query)
	if err != nil {
		return nil, fmt.Errorf("pgxscan GetAllPosts err: %w", err)
	} else if len(postAll) == 0 {
		return nil, models.ErrNotFound
	}

	return postAll, nil
}
