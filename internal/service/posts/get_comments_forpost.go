package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *PostService) GetCommentsForPost(ctx context.Context, post *model.Post, limit, offset int) ([]*model.Comment, error) {
	commentsResp, err := p.Repo.GetCommentsForPost(ctx, post, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("service GetCommentsForPost err: %w", err)
	}

	res := make([]*model.Comment, len(commentsResp))

	for i, comment := range commentsResp {
		res[i] = &model.Comment{
			ID:      comment.ID,
			Author:  comment.Author,
			Content: comment.Content,
			Post:    comment.Post,
			ReplyTo: comment.ReplyTo,
		}
	}
	return res, nil
}
