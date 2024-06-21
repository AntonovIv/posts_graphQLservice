package graph

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, post model.CreatePostReq) (*model.PostListEl, error) {
	newPost, err := r.postService.CreatePost(ctx, post)
	return &newPost, err
}
