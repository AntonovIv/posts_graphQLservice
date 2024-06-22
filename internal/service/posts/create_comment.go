package posts

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) CreateComment(ctx context.Context, input model.CreateCommentReq) (*model.Comment, error) {
	commentResp, err := p.repo.CreateComment(ctx, input)
	return &commentResp, err
}
