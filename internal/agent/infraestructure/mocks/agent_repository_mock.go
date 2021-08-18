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

func (ag *AgentRepositoryMock) Update(ctx context.Context, agent domain.Agent) error {
	ret := ag.Called(ctx, agent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Agent) error); ok {
		r0 = rf(ctx, agent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (ag *AgentRepositoryMock) SearchByCodeAndPassword(ctx context.Context, agentCode int, password int) (domain.Agent, error) {
	ret := ag.Called(ctx, agentCode, password)

	var r0 domain.Agent
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) (domain.Agent, error)); ok {
		r0, r1 = rf(ctx, agentCode, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (ag *AgentRepositoryMock) GetPassword(ctx context.Context, phoneNumber int) (int, error) {
	ret := ag.Called(ctx, phoneNumber)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (int, error)); ok {
		r0, r1 = rf(ctx, phoneNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
