package graph

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

// CommentSubscription is the resolver for the CommentSubscription field.
func (r *subscriptionResolver) CommentSubscription(ctx context.Context, postID int) (<-chan *model.Comment, error) {

	obsId, ch, _ := r.postService.CreateObserver(postID)
	go func(obsId, postId int) {
		<-ctx.Done()
		_ = r.postService.DeleteObserver(postID, obsId)
	}(obsId, postID)

	return ch, nil
}
