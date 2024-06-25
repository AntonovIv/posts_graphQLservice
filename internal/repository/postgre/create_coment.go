package postgre

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (r *repository) CreateComment(ctx context.Context, input model.CreateCommentReq) (model.Comment, error) {
	query := `insert into Comments (content, author, postid, reply_to)
	values ($1, $2, $3, $4)`

	_, err := r.db.DB(ctx).Exec(ctx, query,
		input.Content, input.Author, input.Post, input.ReplyTo)
	if err != nil {
		return model.Comment{}, fmt.Errorf("repository CreateComment err: %w", err)

	}

	return model.Comment{
		Post:    input.Post,
		Author:  input.Author,
		Content: input.Content,
		ReplyTo: input.ReplyTo,
	}, nil
}
