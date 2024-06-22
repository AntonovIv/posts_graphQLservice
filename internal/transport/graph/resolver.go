package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"log/slog"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

type postsSrv interface {
	//post metods
	CreatePost(context.Context, model.CreatePostReq) (*model.PostListEl, error)
	GetPostByID(context.Context, int) (*model.Post, error)
	GetAllPosts(context.Context) ([]*model.PostListEl, error)

	//comment methods
	CreateComment(context.Context, model.CreateCommentReq) (*model.Comment, error)
	GetCommentsForPost(context.Context, *model.Post, int, int) ([]*model.Comment, error)
	GetRepliesComments(context.Context, *model.Comment, int, int) ([]*model.Comment, error)
}

type Resolver struct {
	postService postsSrv
	logger      *slog.Logger
}

func New(postService postsSrv, logger *slog.Logger) *Resolver {
	srv := &Resolver{
		postService: postService,
		logger:      logger,
	}
	return srv
}
