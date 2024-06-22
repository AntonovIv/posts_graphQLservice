package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) GetRepliesComments(ctx context.Context, obj *model.Comment, limit, offset int) ([]model.Comment, error) {
	query := `select id, author, content, postid, reply_to from comments
		where reply_to = $1`

	args := []interface{}{obj.ID}

	if limit > 0 && offset >= 0 {
		query += ` offset $2
		 limit $3`
		args = append(args, offset)
		args = append(args, limit)
	}
	var rsComments []model.Comment

	err := pgxscan.Select(ctx, r.db.DB(ctx), &rsComments, query, args...)
	if err != nil {
		return nil, fmt.Errorf("pgxscan GetRepliesComments err: %w", err)
	}

	return rsComments, nil

}
