package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) GetCommentsForPost(ctx context.Context, obj *model.Post) ([]model.Comment, error) {
	query := `select id, author, content, postid, reply_to from comments
		where postid = $1 
		and reply_to is null`

	var comments []model.Comment

	err := pgxscan.Select(ctx, r.db.DB(ctx), &comments, query, obj.ID)

	if pgxscan.NotFound(err) {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("pgxscan get: %w", err)
	}
	fmt.Println(comments)
	return comments, nil

}
