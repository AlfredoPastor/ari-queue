package domain

import "ari-queue/internal/shared/uniqueids/domain/uuid"

type State string

const (
	WAITING    State = "WAITING"
	CONNECTING State = "CONNECTING"
	ONCALL     State = "ONCALL"
	HANGUP     State = "HANGUP"
)

type Customer struct {
	ID             uuid.VoId
	Name           string
	Number         string
	PositionNumber int64
	State          State
}
