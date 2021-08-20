package add_agent_to_queue

import (
	domainCustomer "ari-queue/internal/customer/domain"
	"ari-queue/internal/queue/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)

type AddCustomerToQueueService struct {
	domain.QueueRepository
	domainCustomer.CustomerRepository
}

func (a AddCustomerToQueueService) Add(ctx context.Context, idQueue uuid.VoId, idCustomer uuid.VoId) error {
	queue, err := a.QueueRepository.Search(ctx, idQueue)
	if err != nil {
		return err
	}
	_, err = a.CustomerRepository.Search(ctx, idQueue)
	if err != nil {
		return err
	}
	queue.AddCustomer(idCustomer)
	err = a.QueueRepository.Upsert(ctx, queue)
	if err != nil {
		return err
	}
	return nil
}

func NewAddCustomerToQueueService(qrepo domain.QueueRepository, crepo domainCustomer.CustomerRepository) AddCustomerToQueueService {
	return AddCustomerToQueueService{
		QueueRepository:    qrepo,
		CustomerRepository: crepo,
	}
}
