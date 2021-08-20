package domain

import (
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)
type CustomerRepository interface {
	Upsert(context.Context, Customer) error
	Delete(context.Context, uuid.VoId) error
	Search(context.Context, uuid.VoId) (Customer, error)
}

type Customer struct {
	ID             uuid.VoId
	Name           string
	Number         string
	PositionNumber int64
	State          VoState
}
