package subscription

func (o *ObserverPoolImpl) DeleteObserver(postId, obsId int) error {
	o.Lock()
	defer o.Unlock()

	obs, ok := o.observers[postId]
	if !ok {
		for i, observer := range obs {
			if observer.id == obsId {
				o.observers[postId] = append(obs[:i], obs[i+1:]...)
			}
		}
	}
	return nil
}
