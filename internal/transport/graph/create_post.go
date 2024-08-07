package graph

import (
	"context"
	"log/slog"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/AntonovIv/post_graphQlservice/internal/validation"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, post model.CreatePostReq) (*model.PostListEl, error) {
	r.logger.DebugContext(ctx, "Create post request")

	if err := validation.CreatePostValidate(post); err != nil {
		r.logger.InfoContext(ctx, "create post request: err",
			slog.Any("err", err))

		return nil, &gqlerror.Error{
			Message: "bad request: " + err.Error(),
		}
	}

	postResp, err := r.postService.CreatePost(ctx, post)
	if err != nil {
		r.logger.ErrorContext(ctx, "Create post request inernal err",
			slog.Any("err", err))

		return nil, models.ErrInternalServerResolver
	}

	return postResp, nil
}
