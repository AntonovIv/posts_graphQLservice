package graph

import (
	"context"
	"log/slog"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, post model.CreatePostReq) (*model.PostListEl, error) {
	r.logger.DebugContext(ctx, "Create post request")

	postResp, err := r.postService.CreatePost(ctx, post)
	if err != nil {
		r.logger.ErrorContext(ctx, "Create post request inernal err",
			slog.Any("err", err))

		return nil, &gqlerror.Error{
			Message: "internal server error",
		}
	}

	return postResp, nil
}
