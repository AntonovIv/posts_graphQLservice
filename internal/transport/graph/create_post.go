package graph

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, post model.CreatePostReq) (*model.PostListEl, error) {
	postResp, err := r.postService.CreatePost(ctx, post)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: "internal server error",
		}
	}

	return postResp, nil
}
