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
