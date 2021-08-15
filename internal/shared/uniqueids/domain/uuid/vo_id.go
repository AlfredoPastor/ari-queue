package uuid

import (
	"ari-queue/internal/shared/uniqueids/domain/tools"
	"errors"
)

type VoId string

var ErrorInvalidID = errors.New("invalid id")

func NewVoId() VoId {
	return VoId(tools.NewUuid())
}

func NewVoIdFromString(value string) (VoId, error) {
	_, err := tools.NewUuidFromString(value)
	if err != nil {
		return VoId(""), ErrorInvalidID
	}
	return VoId(value), nil
}

func (v VoId) String() string {
	return string(v)
}

func (v VoId) Primitive() string {
	return string(v)
}
