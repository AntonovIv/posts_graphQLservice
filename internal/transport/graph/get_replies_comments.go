package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	rsCommentsResp, err := r.postService.GetRepliesComments(ctx, obj)
	return rsCommentsResp, err
}
