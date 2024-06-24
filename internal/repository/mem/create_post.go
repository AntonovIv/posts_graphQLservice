package mem

import (
	"context"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

func (r *MemDb) CreatePost(ctx context.Context, post model.CreatePostReq) (model.PostListEl, error) {
	r.p.Lock()
	defer r.p.Unlock()

	r.p.postCounter++

	hpost := model.Post{
		ID:              r.p.postCounter,
		Name:            post.Name,
		Author:          post.Author,
		Content:         post.Content,
		CommentsAllowed: post.CommentsAllowed,
	}
	r.p.posts = append(r.p.posts, hpost)

	return model.PostListEl{ID: hpost.ID,
		Name:    hpost.Name,
		Author:  hpost.Author,
		Content: hpost.Content}, nil
}
