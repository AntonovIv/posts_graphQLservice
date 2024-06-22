package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) GetCommentsForPost(ctx context.Context, obj *model.Post) ([]model.Comment, error) {
	query := `select id, author, content, postid, reply_to from comments
		where postid = $1 
		and reply_to is null`

	var comments []model.Comment

	err := pgxscan.Select(ctx, r.db.DB(ctx), &comments, query, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("pgxscan get comments for post err: %w", err)
	}

	return comments, nil

}
