package mem

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
)

func (r *MemDb) GetCommentsForPost(ctx context.Context, post *model.Post, limit, offset int) ([]model.Comment, error) {
	r.c.RLock()
	defer r.c.RUnlock()

	var res []model.Comment

	if offset >= len(r.c.comments) {
		return nil, models.ErrNotFound
	}

	for _, comment := range r.c.comments {
		if comment.ReplyTo == nil && comment.Post == post.ID {
			res = append(res, comment)
		}
	}
	if len(res) == 0 || offset >= len(res) {
		return nil, models.ErrNotFound
	}

	if offset+limit > len(res) {
		hres := make([]model.Comment, len(res[offset:]))
		_ = copy(hres, res[offset:])
		return hres, nil
	}

	hres := make([]model.Comment, len(res[offset:(offset)+limit]))
	_ = copy(hres, res[offset:(offset)+limit])

	return hres, nil
}
