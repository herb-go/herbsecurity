package credential

import "github.com/herb-go/herbsecurity/authority"

type Authenticator interface {
	Authenticate(Credentials) (*authority.Auth, error)
	DependencesData() (map[Name]bool, error)
}
type PlainAuthenticator struct {
	authenticateFunc func(Credentials) (*authority.Auth, error)
	dependences      map[Name]bool
}

func (a PlainAuthenticator) Authenticate(c Credentials) (*authority.Auth, error) {
	return a.authenticateFunc(c)
}

func (a PlainAuthenticator) DependencesData() (map[Name]bool, error) {
	return a.dependences, nil
}
func AuthenticatorFunc(a func(Credentials) (*authority.Auth, error), c ...Name) Authenticator {
	authenicator := &PlainAuthenticator{
		authenticateFunc: a,
		dependences:      map[Name]bool{},
	}
	for k := range c {
		authenicator.dependences[c[k]] = true
	}
	return authenicator
}

func Authenticate(a Authenticator, c ...CredentialSource) (*authority.Auth, error) {
	m := NewMap()
	availableTypes, err := a.DependencesData()
	if err != nil {
		return nil, err
	}
	for k := range c {
		n, err := c[k].NameData()
		if err != nil {
			return nil, err
		}
		if availableTypes[n] {
			val, err := c[k].ValueData()
			if err != nil {
				return nil, err
			}
			m.Set(n, val)
		}
	}
	return a.Authenticate(m)
}

type FixedAuthenticator string

func (v FixedAuthenticator) Authenticate(Credentials) (*authority.Auth, error) {
	return authority.NewAuth(authority.Principal(v)), nil
}
func (v FixedAuthenticator) DependencesData() (map[Name]bool, error) {
	return map[Name]bool{}, nil
}

var ForbiddenAuthenticator = AuthenticatorFunc(func(Credentials) (*authority.Auth, error) {
	return nil, nil
})
