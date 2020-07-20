package secret

type Secret []byte

func (s Secret) LoadSecret() (Secret, error) {
	return s, nil
}

type SecretLoader interface {
	LoadSecret() (Secret, error)
}
