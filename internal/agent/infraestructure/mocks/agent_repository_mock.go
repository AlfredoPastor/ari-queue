package mocks

import (
	"ari-queue/internal/agent/domain"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"

	"github.com/stretchr/testify/mock"
)

type AgentRepositoryMock struct {
	mock.Mock
}

func (ag *AgentRepositoryMock) Save(ctx context.Context, agent domain.Agent) error {
	ret := ag.Called(ctx, agent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Agent) error); ok {
		r0 = rf(ctx, agent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (ag *AgentRepositoryMock) Delete(ctx context.Context, id uuid.VoId) error {
	ret := ag.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.VoId) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
