package creator

import (
	"ari-queue/internal/agent/domain"
	"ari-queue/internal/agent/infraestructure/mocks"
	"ari-queue/internal/shared/uniqueids/domain/tools"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Creator_Agent_OK(t *testing.T) {
	id, idPhone, name, state, statePhone := "5efbbb993a65a00001ad8ffd", "5efbbb993a65a00001ad8ffd", "Alfredo", "UNAVAILABLE", "UNAVAILABLE"
	code, phoneNumber := 12, 301
	phone, err := domain.NewPhone(idPhone, phoneNumber, statePhone)
	require.NoError(t, err)

	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("Upsert", mock.Anything, mock.AnythingOfType("domain.Agent")).Return(nil)
	creatorAgentService := NewCreatorAgentService(agentRepositoryMock)
	err = creatorAgentService.Create(context.Background(), id, code, name, state, phone)
	agentRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_Creator_Agent_Upsert_Fail(t *testing.T) {
	id, idPhone, name, state, statePhone := tools.NewUuid(), tools.NewUuid(), "Alfredo", "UNAVAILABLE", "UNAVAILABLE"
	code, phoneNumber := 12, 301
	phone, err := domain.NewPhone(idPhone, phoneNumber, statePhone)
	require.NoError(t, err)

	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("Upsert", mock.Anything, mock.AnythingOfType("domain.Agent")).Return(errors.New("something unexpected happened"))
	creatorAgentService := NewCreatorAgentService(agentRepositoryMock)
	err = creatorAgentService.Create(context.Background(), id, code, name, state, phone)
	agentRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Creator_Agent_ID_Fail(t *testing.T) {
	id, idPhone, name, state, statePhone := "5efbbb993a65a00", tools.NewUuid(), "Alfredo", "UNAVAILABLE", "UNAVAILABLE"
	code, phoneNumber := 12, 301
	phone, err := domain.NewPhone(idPhone, phoneNumber, statePhone)
	require.NoError(t, err)

	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	creatorAgentService := NewCreatorAgentService(agentRepositoryMock)
	err = creatorAgentService.Create(context.Background(), id, code, name, state, phone)
	assert.Error(t, err)
}
