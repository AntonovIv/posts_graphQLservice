package posts

import (
	serviceTm "github.com/AntonovIv/post_graphQlservice/internal/service"
)

type postRepo interface {
	serviceTm.TransactionManager

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
		// repo: repo,
	}
}
