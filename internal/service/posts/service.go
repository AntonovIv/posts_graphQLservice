package posts

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	serviceTm "github.com/AntonovIv/post_graphQlservice/internal/service"
	"github.com/AntonovIv/post_graphQlservice/internal/service/posts/subscription"
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
type ObserverPool interface {
	CreateObserver(postId int) (int, chan *model.Comment, error)
	DeleteObserver(postId, chanId int) error
	NotifyObservers(postId int, comment model.Comment) error
}

type postService struct {
	repo    postRepo
	obsPool ObserverPool
}

func New(repo postRepo) *postService {
	return &postService{
		repo:    repo,
		obsPool: subscription.NewObserverPool(),
	}
}
