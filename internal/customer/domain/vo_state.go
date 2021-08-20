package domain

import "errors"

type VoState string

var ErrStateNotAllow = errors.New("state not allow")

const (
	WAITING    VoState = "WAITING"
	CONNECTING VoState = "CONNECTING"
	ONCALL     VoState = "ONCALL"
	HANGUP     VoState = "HANGUP"
)

func NewVoState(state string) (VoState, error) {
	switch VoState(state) {
	case WAITING:
		return VoState(state), nil
	case CONNECTING:
		return VoState(state), nil
	case ONCALL:
		return VoState(state), nil
	case HANGUP:
		return VoState(state), nil
	default:
		return VoState(""), ErrStateNotAllow
	}
}
