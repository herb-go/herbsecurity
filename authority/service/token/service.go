package token

import (
	"github.com/herb-go/herbsecurity/authority"
)

type Service interface {
	Issuer
	Revoker
	Refresher
	Loader
	Start() error
	Stop() error
}
type Issuer interface {
	IssueToken(*Token) error
}

type Revoker interface {
	RevokeToken(authority.Principal, authority.Passphrase) (Token, error)
}

type Refresher interface {
	RefreshToken(authority.Principal, authority.Passphrase, *authority.Expiration) error
}

type Loader interface {
	LoadToken(authority.Passphrase) (*Token, error)
}

type Manager interface {
	Service
	VerifyToken(authority.Principal, authority.Passphrase) (bool, error)
	ListTokens(authority.Principal) ([]*Token, error)
	RevokeTokens(authority.Principal) error
	GeneratePassphrase() (authority.Passphrase, error)
}
