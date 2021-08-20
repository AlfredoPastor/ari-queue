package signin

import (
	"ari-queue/internal/agent/domain"
	"context"
)

type SigninService struct {
	domain.AgentRepository
}

func (s SigninService) Signin(ctx context.Context, agentCode int, phoneNumber int) error {
	password, err := s.AgentRepository.GetPassword(ctx, phoneNumber)
	if err != nil {
		return err
	}
	agent, err := s.AgentRepository.SearchByCodeAndPassword(ctx, agentCode, password)
	if err != nil {
		return err
	}
	agent.Signin()
	err = s.AgentRepository.Upsert(ctx, agent)
	if err != nil {
		return err
	}
	return nil
}

func NewSigninService(repo domain.AgentRepository) SigninService {
	return SigninService{
		AgentRepository: repo,
	}
}
