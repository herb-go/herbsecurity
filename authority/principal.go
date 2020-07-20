package authority

type Principal string

func (p Principal) PrincipalData() (Principal, error) {
	return p, nil
}

type PrincipalSource interface {
	PrincipalData() (Principal, error)
}
