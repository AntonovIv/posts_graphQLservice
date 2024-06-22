package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) GetAllPosts(ctx context.Context) ([]model.PostListEl, error) {

	query := `select id, name, author, content 
	from posts`

	var postAll []model.PostListEl

	// err := pgxscan.Get(ctx, r.db.DB(ctx), &postAll, query)

	err := pgxscan.Select(ctx, r.db.DB(ctx), &postAll, query)
	if err != nil {
		return nil, fmt.Errorf("pgxscan get: %w", err)
	}

	fmt.Println(postAll)
	return postAll, nil
}
