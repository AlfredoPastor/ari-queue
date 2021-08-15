package creator_agent

import (
	"ari-queue/internal/agent/infraestructure/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreatorAgent(t *testing.T) {
	id, name := "78t78g8878787g78g87g", "Alfredo"
	code, phoneNumber := 12, 32
	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Agent")).Return(errors.New("something unexpected happened"))
	creatorAgentService := NewCreatorAgentService(agentRepositoryMock)
	err := creatorAgentService.Create(context.TODO(), id, code, phoneNumber, name)
	agentRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}
