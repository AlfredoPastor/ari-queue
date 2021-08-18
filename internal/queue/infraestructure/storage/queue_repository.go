package storage

import (
	"ari-queue/internal/queue/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
	"errors"
	"sync"
)

type QueueRepository struct {
	sync.Mutex
	queues map[uuid.VoId]domain.Queue
}

func (q *QueueRepository) Upsert(ctx context.Context, queue domain.Queue) error {
	q.Lock()
	defer q.Unlock()
	q.queues[queue.ID] = queue
	return nil
}

func (q *QueueRepository) Delete(ctx context.Context, id uuid.VoId) error {
	q.Lock()
	defer q.Unlock()
	delete(q.queues, id)
	return nil
}

func (q *QueueRepository) Search(ctx context.Context, id uuid.VoId) (domain.Queue, error) {
	q.Lock()
	defer q.Unlock()
	queue, ok := q.queues[id]
	if !ok {
		return domain.Queue{}, errors.New("queue not found")
	}
	return queue, nil
}

func NewQueueRepository() QueueRepository {
	return QueueRepository{
		queues: make(map[uuid.VoId]domain.Queue),
	}
}
