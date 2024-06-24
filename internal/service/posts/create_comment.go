package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) CreateComment(ctx context.Context, input model.CreateCommentReq) (*model.Comment, error) {

	var commentResp model.Comment

	err := p.repo.WithTransaction(ctx, func(ctx context.Context) error {
		post, err := p.repo.GetPostByID(ctx, input.Post)
		if err != nil {
			return fmt.Errorf("service create comment func failed get: %w", err)
		}
		if !post.CommentsAllowed {
			return fmt.Errorf("comments are not allowed for post: %d", post.ID)
		}

		commentResp, err = p.repo.CreateComment(ctx, input)
		if err != nil {
			return fmt.Errorf("service create comment failed create: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &commentResp, nil
}
