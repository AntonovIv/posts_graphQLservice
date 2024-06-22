package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

type postsSrv interface {
	//post metods
	CreatePost(context.Context, model.CreatePostReq) (*model.PostListEl, error)
	GetPostByID(context.Context, int) (*model.Post, error)
	GetAllPosts(context.Context) ([]*model.PostListEl, error)

	//comment methods
	CreateComment(context.Context, model.CreateCommentReq) (*model.Comment, error)
	GetCommentsForPost(context.Context, *model.Post) ([]*model.Comment, error)
	GetRepliesComments(context.Context, *model.Comment) ([]*model.Comment, error)
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
