package token

import (
	"github.com/herb-go/herbsecurity/authority"
)

type Token struct {
	authority.Passphrase
	authority.Principal
	authority.Agent
	*authority.Expiration
	Payloads authority.Payloads
}

func (t *Token) Auth() *authority.Auth {
	return authority.NewAuth(t.Principal).WithAgent(t.Agent).WithPayloads(t.Payloads)
}

func NewToken() *Token {
	return &Token{}
}
