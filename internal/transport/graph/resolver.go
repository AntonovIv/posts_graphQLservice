package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

type postsSrv interface {
	CreatePost(context.Context, model.CreatePostReq) (model.PostListEl, error)
	//GetPostById(id int) (model.Post, error)
	// GetAllPosts(page, pageSize *int) ([]models.Post, error)
}

type Resolver struct {
	postService postsSrv
}

func New(postService postsSrv) *Resolver {
	srv := &Resolver{
		postService: postService,
	}
	return srv
}
