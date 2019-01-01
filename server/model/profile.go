package model

import "github.com/satori/go.uuid"

func CreateProfile(clientID string) *Profile {
	return &Profile{ClientID: clientID, Secret: uuid.NewV4().String()}
}

type Profile struct {
	ClientID string
	Nick     string
	Secret   string
}

func (p *Profile) IsMe(secret string) bool {
	return p.Secret == secret
}
