package application

import (
	"ari-queue/internal/agent/infraestructure/mocks"
	"ari-queue/internal/shared/uniqueids/domain/tools"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Deleter_Agent_OK(t *testing.T) {
	id := tools.NewUuid()
	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("Delete", mock.Anything, mock.AnythingOfType("uuid.VoId")).Return(nil)
	Service := NewDeleterAgentService(agentRepositoryMock)
	err := Service.Delete(context.Background(), uuid.VoId(id))
	agentRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_Deleter_Agent_Fail(t *testing.T) {
	id := tools.NewUuid()
	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("Delete", mock.Anything, mock.AnythingOfType("uuid.VoId")).Return(errors.New("upsss"))
	Service := NewDeleterAgentService(agentRepositoryMock)
	err := Service.Delete(context.Background(), uuid.VoId(id))
	agentRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}
