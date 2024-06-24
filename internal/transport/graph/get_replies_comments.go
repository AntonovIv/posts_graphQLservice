package graph

// // This file will be automatically regenerated based on the schema, any resolver implementations
// // will be copied through when generating and any unknown code will be moved to the end.
// // Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"log/slog"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/validation"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment, limit *int, offset *int) ([]*model.Comment, error) {
	r.logger.DebugContext(ctx, "get replydComments request")

	lim, offs := validation.PagingValidate(limit, offset)
	rsCommentsResp, err := r.postService.GetRepliesComments(ctx, obj, lim, offs)
	if err != nil {
		r.logger.ErrorContext(ctx, "get replyComments request: err",
			slog.Any("err", err))

		return nil, &gqlerror.Error{
			Message: "internal server error",
		}
	}

	return rsCommentsResp, nil
}
