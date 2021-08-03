package domain

import "github.com/AlfredoPastor/ari-queue/internal/shared/uniqueids/domain/uuid"

type Agent struct {
	ID     uuid.VoId
	Number int
	Phone  int
}
