package domain

import (
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
)

//go:generate mockery --outpkg=mocks --output=../../infraestructure/mocks --name=AgentRepository
type AgentRepository interface {
	GetPassword(ctx context.Context, phoneNumber int) (int, error)
	SearchByCodeAndPassword(ctx context.Context, agentCode int, password int) (Agent, error)
	Save(context.Context, Agent) error
	Update(context.Context, Agent) error
	Delete(context.Context, uuid.VoId) error
}
type Agent struct {
	ID    uuid.VoId
	Code  int
	Name  string
	State VoAgentState
	Phone
}

func NewAgent(id string, code int, name, state string, phone Phone) (Agent, error) {
	idVo, err := uuid.NewVoIdFromString(id)
	if err != nil {
		return Agent{}, err
	}
	stateVo := NewVoAgentState(state)
	return Agent{
		ID:    idVo,
		Code:  code,
		Name:  name,
		State: stateVo,
		Phone: phone,
	}, nil
}

func (a Agent) Signin() {
	a.State = SIGNIN
}
