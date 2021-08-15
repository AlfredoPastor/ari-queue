package domain

import "ari-queue/internal/shared/uniqueids/domain/uuid"

type Phone struct {
	ID     uuid.VoId
	State  VoPhoneState
	Number VoPhoneNumber
}

func NewPhone(id string, number int, state string) (Phone, error) {
	idVo, err := uuid.NewVoIdFromString(id)
	if err != nil {
		return Phone{}, err
	}
	numberVo := NewVoPhoneNumber(number)
	stateVo := NewVoPhoneState(state)
	return Phone{
		ID:     idVo,
		Number: numberVo,
		State:  stateVo,
	}, nil
}
