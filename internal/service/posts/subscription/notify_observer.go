package subscription

import "github.com/AntonovIv/post_graphQlservice/graph/model"

func (c *ObserverPoolImpl) NotifyObservers(postId int, comment model.Comment) error {
	c.Lock()
	defer c.Unlock()

	obs, ok := c.observers[postId]
	if ok {
		for _, observer := range obs {
			observer.ch <- &comment
		}
	}
	return nil
}
