package application

type CreatorQueue struct{}

func (c CreatorQueue) Create() error {
	return nil
}

func NewCreatorQueue() CreatorQueue {
	return CreatorQueue{}
}
