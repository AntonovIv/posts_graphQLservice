package mem

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"

	"github.com/AntonovIv/post_graphQlservice/internal/models"
)

func (r *MemDb) GetPostByID(ctx context.Context, id int) (model.Post, error) {
	r.p.RLock()
	defer r.p.RUnlock()

	if id > r.p.postCounter {
		return model.Post{}, models.ErrNotFound
	}

	return r.p.posts[id-1], nil
}
