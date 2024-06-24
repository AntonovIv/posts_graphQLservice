package postgre

import (
	"context"
	"fmt"
	"log"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *repository) GetCommentsForPost(ctx context.Context, obj *model.Post, limit, offset int) ([]model.Comment, error) {
	query := `select id, author, content, postid, reply_to from comments
		where postid = $1 
		and reply_to is null`

	args := []interface{}{obj.ID}

	if limit > 0 && offset >= 0 {
		query += ` offset $2
		 limit $3`
		args = append(args, offset)
		args = append(args, limit)
	}
	var comments []model.Comment

	log.Println("\n", query)

	log.Println("\n", obj.ID)
	err := pgxscan.Select(ctx, r.db.DB(ctx), &comments, query, args...)
	if err != nil {
		return nil, fmt.Errorf("pgxscan GetCommentsForPost err: %w", err)
	}

	return comments, nil

}
