package mem

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (r *MemDb) CreateComment(ctx context.Context, input model.CreateCommentReq) (model.Comment, error) {
	r.c.Lock()
	defer r.c.Unlock()

	r.c.commCounter++

	r.c.comments = append(r.c.comments, model.Comment{
		ID:      r.c.commCounter,
		Post:    input.Post,
		Author:  input.Author,
		Content: input.Content,
		ReplyTo: input.ReplyTo,
	})

	return r.c.comments[r.c.commCounter-1], nil

}
