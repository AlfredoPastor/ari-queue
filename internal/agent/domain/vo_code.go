package domain

type VoCode int

func NewVoCode(number int) VoCode {
	return VoCode(number)
}
