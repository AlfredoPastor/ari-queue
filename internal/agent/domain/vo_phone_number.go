package domain

type VoPhoneNumber int

func NewVoPhoneNumber(number int) VoPhoneNumber {
	return VoPhoneNumber(number)
}
