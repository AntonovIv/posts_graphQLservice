package mem

import (
	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (r *MemDb) GetAllPosts() ([]model.Post, error) {
	r.p.RLock()
	defer r.p.RUnlock()

	return r.p.posts, nil
}
