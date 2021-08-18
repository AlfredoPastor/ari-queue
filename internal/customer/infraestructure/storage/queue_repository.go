package storage

import (
	"ari-queue/internal/customer/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
	"errors"
	"sync"
)

type CustomerRepository struct {
	sync.Mutex
	customers map[uuid.VoId]domain.Customer
}

func (q *CustomerRepository) Upsert(ctx context.Context, customer domain.Customer) error {
	q.Lock()
	defer q.Unlock()
	q.customers[customer.ID] = customer
	return nil
}

func (q *CustomerRepository) Delete(ctx context.Context, id uuid.VoId) error {
	q.Lock()
	defer q.Unlock()
	delete(q.customers, id)
	return nil
}

func (q *CustomerRepository) Search(ctx context.Context, id uuid.VoId) (domain.Customer, error) {
	q.Lock()
	defer q.Unlock()
	customer, ok := q.customers[id]
	if !ok {
		return domain.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

func NewCustomerRepository() CustomerRepository {
	return CustomerRepository{
		customers: make(map[uuid.VoId]domain.Customer),
	}
}
