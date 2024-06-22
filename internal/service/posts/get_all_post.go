package posts

import (
	"context"
	"fmt"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) GetAllPosts(ctx context.Context) ([]*model.PostListEl, error) {
	postAllResp, err := p.repo.GetAllPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("service GetAllPosts err: %w", err)
	}

	res := make([]*model.PostListEl, len(postAllResp))

	for i, post := range postAllResp {
		res[i] = &model.PostListEl{
			ID:      post.ID,
			Name:    post.Name,
			Author:  post.Author,
			Content: post.Content,
		}
	}

	return res, nil
}
