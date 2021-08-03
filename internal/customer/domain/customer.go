package domain

import "github.com/AlfredoPastor/ari-queue/internal/shared/uniqueids/domain/uuid"

type Customer struct {
	ID     uuid.VoId
	Name   string
	Number string
}
