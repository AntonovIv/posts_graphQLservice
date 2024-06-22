package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) GetRepliesComments(ctx context.Context, obj *model.Comment) ([]model.Comment, error) {
	query := `select id, author, content, postid, reply_to from comments
		where reply_to = $1`

	var rsComments []model.Comment

	err := pgxscan.Select(ctx, r.db.DB(ctx), &rsComments, query, obj.ID)
	if err != nil {
		return nil, fmt.Errorf("pgxscan GetRepliesComments err: %w", err)
	}

	return rsComments, nil

}
