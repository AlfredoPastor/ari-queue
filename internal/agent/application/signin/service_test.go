package signin

import (
	"ari-queue/internal/agent/domain"
	"ari-queue/internal/agent/infraestructure/mocks"
	"ari-queue/internal/shared/uniqueids/domain/tools"
	"ari-queue/internal/shared/uniqueids/domain/uuid"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Signin_Ok(t *testing.T) {
	agentCode, phoneNumber := 344, 301
	agent := domain.Agent{
		ID:   uuid.VoId(tools.NewUuid()),
		Code: agentCode,
		Name: "Alfredo",
		Phone: domain.Phone{
			ID:     uuid.VoId(tools.NewUuid()),
			Number: domain.VoPhoneNumber(phoneNumber),
		},
	}
	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("GetPassword", mock.Anything, mock.AnythingOfType("int")).Return(777, nil)
	agentRepositoryMock.On("SearchByCodeAndPassword", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(agent, nil)
	agentRepositoryMock.On("Upsert", mock.Anything, mock.AnythingOfType("domain.Agent")).Return(nil)
	siginService := NewSigninService(agentRepositoryMock)
	err := siginService.Signin(context.Background(), agentCode, phoneNumber)
	agentRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_Signin_GetPassword_Fail(t *testing.T) {
	agentCode, phoneNumber := 344, 301

	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("GetPassword", mock.Anything, mock.AnythingOfType("int")).Return(777, errors.New("upsss"))
	siginService := NewSigninService(agentRepositoryMock)
	err := siginService.Signin(context.Background(), agentCode, phoneNumber)
	agentRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Signin_SearchByCodeAndPassword_Fail(t *testing.T) {
	agentCode, phoneNumber := 344, 301

	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("GetPassword", mock.Anything, mock.AnythingOfType("int")).Return(777, nil)
	agentRepositoryMock.On("SearchByCodeAndPassword", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain.Agent{}, errors.New("upssss"))
	siginService := NewSigninService(agentRepositoryMock)
	err := siginService.Signin(context.Background(), agentCode, phoneNumber)
	agentRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Signin_Upsert_Fail(t *testing.T) {
	agentCode, phoneNumber := 344, 301
	agent := domain.Agent{
		ID:   uuid.VoId(tools.NewUuid()),
		Code: agentCode,
		Name: "Alfredo",
		Phone: domain.Phone{
			ID:     uuid.VoId(tools.NewUuid()),
			Number: domain.VoPhoneNumber(phoneNumber),
		},
	}
	agentRepositoryMock := new(mocks.AgentRepositoryMock)
	agentRepositoryMock.On("GetPassword", mock.Anything, mock.AnythingOfType("int")).Return(777, nil)
	agentRepositoryMock.On("SearchByCodeAndPassword", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(agent, nil)
	agentRepositoryMock.On("Upsert", mock.Anything, mock.AnythingOfType("domain.Agent")).Return(errors.New("upsss"))
	siginService := NewSigninService(agentRepositoryMock)
	err := siginService.Signin(context.Background(), agentCode, phoneNumber)
	agentRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}
