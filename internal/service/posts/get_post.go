package posts

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) GetPostByID(ctx context.Context, id int) (*model.Post, error) {
	postResp, err := p.repo.GetPostByID(ctx, id)
	return &postResp, err
}
