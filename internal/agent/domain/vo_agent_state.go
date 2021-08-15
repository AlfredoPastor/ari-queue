package domain

type VoAgentState string

const (
	AVAILABLE VoAgentState = "AVAILABLE"
	PAUSE     VoAgentState = "PAUSE"
	BUSY      VoAgentState = "BUSY"
)

func NewVoAgentState(state string) VoAgentState {
	return VoAgentState(state)
}
