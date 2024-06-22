package posts

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (p *postService) GetAllPosts(ctx context.Context) ([]*model.PostListEl, error) {
	postAllResp, err := p.repo.GetAllPosts(ctx)

	res := make([]*model.PostListEl, len(postAllResp))
	for i, post := range postAllResp {
		res[i] = &model.PostListEl{
			ID:      post.ID,
			Name:    post.Name,
			Author:  post.Author,
			Content: post.Content,
		}
	}

	return res, err
}
