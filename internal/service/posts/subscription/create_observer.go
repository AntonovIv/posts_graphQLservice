package subscription

import "github.com/AntonovIv/post_graphQlservice/graph/model"

func (o *ObserverPoolImpl) CreateObserver(postId int) (int, chan *model.Comment, error) {
	o.Lock()
	defer o.Unlock()

	ch := make(chan *model.Comment)
	o.obsIdCounter++

	o.observers[postId] = append(o.observers[postId], Observer{id: o.obsIdCounter, ch: ch})

	return o.obsIdCounter, ch, nil

}
