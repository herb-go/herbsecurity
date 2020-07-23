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
	*authority.Payloads
	Application
}

func NewVerified() *Verified {
	return &Verified{
		Payloads: authority.NewPayloads(),
	}
}
func (v *Verified) Auth() *authority.Auth {
	if v == nil {
		return nil
	}
	p := v.Payloads.Clone()
	p.Set(authority.PayloadSignSecret, []byte(v.Passphrase))
	return authority.NewAuth(v.Principal).WithAgent(v.Agent).WithPayloads(p)
}
