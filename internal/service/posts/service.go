package posts

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	serviceTm "github.com/AntonovIv/post_graphQlservice/internal/service"
)

type postRepo interface {
	serviceTm.TransactionManager

	CreatePost(context.Context, model.CreatePostReq) (model.PostListEl, error)
	GetPostByID(context.Context, int) (model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.PostListEl, error)
	GetCommentsForPost(context.Context, *model.Post) ([]model.Comment, error)
	GetRepliesComments(context.Context, *model.Comment) ([]model.Comment, error)
	CreateComment(context.Context, model.CreateCommentReq) (model.Comment, error)

	// Get(ctx context.Context, key string) (*model.Record, error)
	// GetAll(ctx context.Context) ([]model.Record, error)
	// Set(ctx context.Context, record model.Record) error
	// Delete(ctx context.Context, key string) error
}

type postService struct {
	repo postRepo
}

func New(repo postRepo) *postService {
	return &postService{
		repo: repo,
	}
}
