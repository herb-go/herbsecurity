package credential

type Name string

func (n Name) NameData() (Name, error) {
	return n, nil
}

type NameSource interface {
	NameData() (Name, error)
}

type Value []byte

func (v Value) ValueData() (Value, error) {
	return v, nil
}

type ValueSource interface {
	ValueData() (Value, error)
}
type Credential struct {
	Name
	Value
}

type CredentialSource interface {
	NameSource
	ValueSource
}

type Credentials interface {
	Get(Name) (Value, error)
	Set(Name, Value) error
}
