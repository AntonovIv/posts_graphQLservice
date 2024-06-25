package graph

import (
	"context"
	"log/slog"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/validation"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CommentSubscription is the resolver for the CommentSubscription field.
func (r *subscriptionResolver) CommentSubscription(ctx context.Context, postID int) (<-chan *model.Comment, error) {
	if err := validation.GetPostValidate(postID); err != nil {
		r.logger.InfoContext(ctx, "cubscription request : err",
			slog.Any("err", err))

		return nil, &gqlerror.Error{
			Message: "bad request: " + err.Error(),
		}
	}

	obsId, ch, _ := r.postService.CreateObserver(postID)
	go func(obsId, postId int) {
		<-ctx.Done()
		_ = r.postService.DeleteObserver(postID, obsId)
	}(obsId, postID)

	return ch, nil
}
