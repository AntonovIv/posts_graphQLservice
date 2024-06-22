package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) CreateComment(ctx context.Context, input model.CreateCommentReq) (*model.Comment, error) {
	commentResp, err := p.repo.CreateComment(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("create CreateComment err: %w", err)
	}

	return &commentResp, nil
}
