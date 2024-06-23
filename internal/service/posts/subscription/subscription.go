package subscription

import (
	"sync"

	"github.com/AntonovIv/post_graphQlservice/graph/model"
)

type ObserverPoolImpl struct {
	obsIdCounter int
	observers    map[int][]Observer
	sync.Mutex
}
type Observer struct {
	id int
	ch chan *model.Comment
}

func NewObserverPool() *ObserverPoolImpl {
	return &ObserverPoolImpl{
		observers: make(map[int][]Observer),
	}
}
