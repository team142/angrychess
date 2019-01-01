package model

func CreateProfile(nick string) *Profile {
	return &Profile{Nick: nick}
}

type Profile struct {
	Nick string
	Out  chan []byte
	In   chan []byte
}
