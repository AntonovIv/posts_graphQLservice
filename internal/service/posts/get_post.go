package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) GetPostByID(ctx context.Context, id int) (*model.Post, error) {
	postResp, err := p.repo.GetPostByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service GetPostByID err: %w", err)
	}
	return &postResp, nil
}