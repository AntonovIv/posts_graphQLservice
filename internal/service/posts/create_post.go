package posts

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) CreatePost(ctx context.Context, post model.CreatePostReq) (*model.PostListEl, error) {
	postResp, err := p.repo.CreatePost(ctx, post)
	return &postResp, err
}
