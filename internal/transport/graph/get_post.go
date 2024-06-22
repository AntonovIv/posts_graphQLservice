package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GetPostByID is the resolver for the GetPostById field.
func (r *queryResolver) GetPostByID(ctx context.Context, id int) (*model.Post, error) {
	postResp, err := r.postService.GetPostByID(ctx, id)
	if errors.Is(err, models.ErrNotFound) {
		return nil, &gqlerror.Error{
			Message: "post not found",
		}
	} else if err != nil {
		return nil, &gqlerror.Error{
			Message: "internal server error",
		}

	}

	return postResp, nil
}
