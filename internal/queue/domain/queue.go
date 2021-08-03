package domain

import (
	"github.com/AlfredoPastor/ari-queue/internal/customer/domain"
	"github.com/AlfredoPastor/ari-queue/internal/shared/uniqueids/domain/uuid"
)

type Queue struct {
	ID        uuid.VoId
	Customers []uuid.VoId
	Agents    []uuid.VoId
}
