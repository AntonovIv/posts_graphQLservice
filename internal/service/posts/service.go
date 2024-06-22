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

	CreateComment(context.Context, model.CreateCommentReq) (model.Comment, error)

	GetCommentsForPost(context.Context, *model.Post, int, int) ([]model.Comment, error)
	GetRepliesComments(context.Context, *model.Comment, int, int) ([]model.Comment, error)
}

type postService struct {
	repo postRepo
}

func New(repo postRepo) *postService {
	return &postService{
		repo: repo,
	}
}
