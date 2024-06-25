package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *PostService) CreatePost(ctx context.Context, post model.CreatePostReq) (*model.PostListEl, error) {
	postResp, err := p.Repo.CreatePost(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("create CreatePost err: %w", err)
	}

	return &postResp, nil
}
