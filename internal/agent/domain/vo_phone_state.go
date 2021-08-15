package domain

type VoPhoneState string

const (
	UNAVAILABLEPHONE VoPhoneState = "UNAVAILABLE"
	AVAILABLEPHONE   VoPhoneState = "AVAILABLE"
)

func NewVoPhoneState(state string) VoPhoneState {
	return VoPhoneState(state)
}
