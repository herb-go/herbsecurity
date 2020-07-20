package authority

type Passphrase string

func (p Passphrase) PassphraseData() (Passphrase, error) {
	return p, nil
}

type PassphraseSource interface {
	PassphraseData() (Passphrase, error)
}
