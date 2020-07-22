package authority

type Principal string

func (p Principal) PrincipalData() (Principal, error) {
	return p, nil
}

func (p Principal) String() string {
	return string(p)
}

type PrincipalSource interface {
	PrincipalData() (Principal, error)
}
