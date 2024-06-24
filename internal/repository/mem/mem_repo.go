package mem

import (
	"sync"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

type MemDb struct {
	p PostsInMemory
	c CommentsInMemory
}
type PostsInMemory struct {
	postCounter int

	posts []model.Post
	sync.RWMutex
}

type CommentsInMemory struct {
	commCounter int

	comments []model.Comment
	sync.RWMutex
}

func newPostsInMemory(size int) PostsInMemory {
	return PostsInMemory{
		postCounter: 0,
		posts:       make([]model.Post, 0, size),
	}
}
func newCommentsInMemory(size int) CommentsInMemory {
	return CommentsInMemory{
		commCounter: 0,
		comments:    make([]model.Comment, 0, size),
	}
}

func NewInMemDb(size int) *MemDb {
	return &MemDb{
		p: newPostsInMemory(size),
		c: newCommentsInMemory(size),
	}
}
