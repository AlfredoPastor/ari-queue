package domain

import "ari-queue/internal/shared/uniqueids/domain/uuid"

type StatePhone string

const (
	UNAVAILABLEPHONE StatePhone = "UNAVAILABLE"
	AVAILABLEPHONE   StatePhone = "AVAILABLE"
)

type Phone struct {
	ID     uuid.VoId
	Number int
	StatePhone
}

func NewPhone(id string, code, number int) Agent {
	return Agent{}
}
