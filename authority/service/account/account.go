package account

import (
	"github.com/herb-go/herbsecurity/authority"
)

type Account struct {
	authority.Authority
	authority.Passphrase
}

type Verified struct {
	*authority.Expiration
	authority.Principal
	authority.Agent
	*authority.Payloads
}

func (v *Verified) Auth() *authority.Auth {
	if v == nil {
		return nil
	}
	return authority.NewAuth(v.Principal).WithAgent(v.Agent)
}

func NewVerified() *Verified {
	return &Verified{
		Payloads: authority.NewPayloads(),
	}
}
