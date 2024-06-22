package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) GetRepliesComments(ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	rsComments, err := p.repo.GetRepliesComments(ctx, obj)
	if err != nil {
		return nil, fmt.Errorf("service GetRepliesComments err:%w", err)
	}
	res := make([]*model.Comment, len(rsComments))

	for i, comment := range rsComments {
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
