package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) CreatePost(ctx context.Context, post model.CreatePostReq) (*model.PostListEl, error) {
	postResp, err := p.repo.CreatePost(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("create post service err: %w", err)
	}

	return &postResp, nil
}
