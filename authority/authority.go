package authority

type Authority string

func (a Authority) AuthorityData() (Authority, error) {
	return a, nil
}
func (a Authority) String() string {
	return string(a)
}

type AuthoritySource interface {
	AuthorityData() (Authority, error)
}
