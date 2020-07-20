package application

import (
	"github.com/herb-go/herbsecurity/authority"
)

type Application struct {
	authority.Authority
	authority.Passphrase
}

type Verified struct {
	authority.Principal
	authority.Agent
	authority.Expiration
	Application
}

func NewVerified() *Verified {
	return &Verified{}
}
func (v *Verified) Auth() *authority.Auth {
	return authority.NewAuth(v.Principal).WithAgent(v.Agent)
}
