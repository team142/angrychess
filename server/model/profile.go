package model

func CreateProfile(nick string) *Profile {
	return &Profile{Nick: nick}
}

type Profile struct {
	Nick   string
	Secret string
	Out    chan []byte
	In     chan []byte
}

func (p *Profile) IsMe(secret string) bool {
	return p.Secret == secret
}
