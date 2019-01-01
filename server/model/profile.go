package model

func CreateProfile(clientID string) *Profile {
	return &Profile{ClientID: clientID}
}

type Profile struct {
	ClientID string
	Nick     string
	Secret   string
}

func (p *Profile) IsMe(secret string) bool {
	return p.Secret == secret
}
