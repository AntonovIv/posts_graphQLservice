package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
)

func (p *PostService) CreateComment(ctx context.Context, input model.CreateCommentReq) (*model.Comment, error) {

	var commentResp model.Comment

	err := p.Repo.WithTransaction(ctx, func(ctx context.Context) error {
		post, err := p.Repo.GetPostByID(ctx, input.Post)
		if err != nil {
			return fmt.Errorf("service create comment func failed get: %w", err)
		}
		if !post.CommentsAllowed {
			return fmt.Errorf("comments are not allowed for post: %w", models.ErrBadPostId)
		}

		commentResp, err = p.Repo.CreateComment(ctx, input)
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
