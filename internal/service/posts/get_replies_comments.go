package posts

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) GetRepliesComments(ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	rsComments, err := p.repo.GetRepliesComments(ctx, obj)

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
