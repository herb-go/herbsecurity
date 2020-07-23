package application

import "github.com/herb-go/herbsecurity/authority"

type Creator interface {
	CreateApplication(authority.Principal, authority.Agent, *authority.Payloads, *authority.Payloads) (*Verified, error)
}

type Regenerator interface {
	RegenerateApplication(authority.Principal, authority.Authority) error
}

type Revoker interface {
	RevokeApplication(authority.Principal, authority.Authority) error
}
type Loader interface {
	LoadApplication(authority.Authority) (*Verified, error)
}

type Service interface {
	Creator
	Regenerator
	Revoker
	Loader
}

type ServiceFactory interface {
	CreateApplicationService(func(interface{}) error) (Service, error)
}
