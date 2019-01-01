package model

import (
	"github.com/satori/go.uuid"
	"github.com/team142/chessfor4/io/ws"
)

func CreateProfile(client *ws.Client) *Profile {
	return &Profile{Client: client, Secret: uuid.NewV4().String()}
}

type Profile struct {
	Client *ws.Client
	Nick   string
	Secret string
}

func (p *Profile) IsMe(secret string) bool {
	return p.Secret == secret
}
