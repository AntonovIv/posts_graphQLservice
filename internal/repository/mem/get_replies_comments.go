package mem

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
	"github.com/AntonovIv/post_graphQlservice/internal/models"
)

func (r *MemDb) GetRepliesComments(ctx context.Context, com *model.Comment, limit, offset int) ([]model.Comment, error) {
	r.c.RLock()
	defer r.c.RUnlock()

	if com.ID > r.c.commCounter || limit*offset > len(r.c.comments)+1 {
		return nil, models.ErrNotFound
	}

	var res []model.Comment

	for _, comment := range r.c.comments {
		if comment.ReplyTo != nil && *comment.ReplyTo == com.ID {
			res = append(res, comment)
		}
	}
	if len(res) == 0 || limit*offset > len(res)+1 {
		return res, nil
	}

	if (limit*offset)+limit > len(res)+1 {
		hres := make([]model.Comment, len(res))
		_ = copy(hres, res[limit*offset:])
		return hres, nil
	}
	//пахнущие костыли исправить
	hres := make([]model.Comment, len(res[limit*offset:(limit*offset)+limit]))
	_ = copy(hres, res[limit*offset:(limit*offset)+limit])

	return hres, nil
}