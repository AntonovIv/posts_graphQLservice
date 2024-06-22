package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"
	"log/slog"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/AntonovIv/post_graphQlservice/internal/validation"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GetPostByID is the resolver for the GetPostById field.
func (r *queryResolver) GetPostByID(ctx context.Context, id int) (*model.Post, error) {
	r.logger.DebugContext(ctx, "getpostbyid request")

	if err := validation.GetPostValidate(id); err != nil {
		r.logger.InfoContext(ctx, "getPostById request: err",
			slog.Any("err", err))

		return nil, &gqlerror.Error{
			Message: "bad request",
		}
	}

	postResp, err := r.postService.GetPostByID(ctx, id)
	if errors.Is(err, models.ErrNotFound) {
		r.logger.ErrorContext(ctx, "getPostById request: err",
			slog.Any("err", err))

		return nil, &gqlerror.Error{
			Message: "post not found",
		}
	} else if err != nil {
		r.logger.ErrorContext(ctx, "getPostById request: err",
			slog.Any("err", err))

		return nil, &gqlerror.Error{
			Message: "internal server error",
		}

	}

	return postResp, nil
}
