package domain

import (
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)

type AgentRepository interface {
	Save(context.Context, Agent) error
	Delete(context.Context, uuid.VoId) error
}

type State string

const (
	AVAILABLE State = "AVAILABLE"
	PAUSE     State = "PAUSE"
	BUSY      State = "BUSY"
)

type Agent struct {
	ID   uuid.VoId
	Code int
	Name string
	State
	Phone
}

func NewAgent(id string, code, phoneNumber int, name string) Agent {
	return Agent{}
}
