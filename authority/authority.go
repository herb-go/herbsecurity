package authority

type Authority string

func (a Authority) AuthorityData() (Authority, error) {
	return a, nil
}

type AuthoritySource interface {
	AuthorityData() (Authority, error)
}
