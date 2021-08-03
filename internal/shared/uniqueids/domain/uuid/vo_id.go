package uuid

import (
	"errors"

	"github.com/AlfredoPastor/ari-queue/internal/shared/uniqueids/domain/tools"
)

type VoId string

var ErrorInvalidID = errors.New("invalid id")

func NewVoId() (VoId, error) {
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
