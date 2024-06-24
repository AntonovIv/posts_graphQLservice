package mem

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (r *MemDb) GetAllPosts(ctx context.Context) ([]model.PostListEl, error) {
	r.p.RLock()
	defer r.p.RUnlock()
	hpostlist := make([]model.PostListEl, len(r.p.posts))
	for i, post := range r.p.posts {
		hpostlist[i] = model.PostListEl{
			ID:      post.ID,
			Name:    post.Name,
			Author:  post.Author,
			Content: post.Content,
		}
	}
	return hpostlist, nil
}
