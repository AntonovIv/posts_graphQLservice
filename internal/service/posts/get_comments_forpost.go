package posts

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) GetCommentsForPost(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	commentsResp, err := p.repo.GetCommentsForPost(ctx, obj)

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
	return res, err
}
