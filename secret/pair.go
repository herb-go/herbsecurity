package secret

type PairLoader interface {
	LoadPublicSecret() (Secret, error)
	LoadPrivateSecret() (Secret, error)
}

type Pair struct {
	Public  Secret
	Private Secret
}

func (p Pair) LoadPublicSecret() (Secret, error) {
	return p.Public, nil
}
func (p Pair) LoadPrivateSecret() (Secret, error) {
	return p.Private, nil
}

func (p Pair) LoadSecret() (Secret, error) {
	return p.Private, nil
}
