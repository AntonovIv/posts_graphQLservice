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

	if limit*offset >= len(r.c.comments) {
		return nil, models.ErrNotFound
	}

	for _, comment := range r.c.comments {
		if comment.ReplyTo == nil && comment.Post == post.ID {
			hres := comment
			res = append(res, hres)
		}
	}
	if len(res) == 0 || limit*offset >= len(res) {
		return nil, models.ErrNotFound
	}

	if (limit*offset)+limit > len(res) {
		hres := make([]model.Comment, len(res))
		_ = copy(hres, res[limit*offset:])
		return hres, nil
	}

	hres := make([]model.Comment, len(res[limit*offset:(limit*offset)+limit]))
	_ = copy(hres, res[limit*offset:(limit*offset)+limit])

	return hres, nil
}
