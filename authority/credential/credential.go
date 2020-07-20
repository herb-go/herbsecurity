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

func (c *Credential) WithName(n Name) *Credential {
	c.Name = n
	return c
}

func (c *Credential) WithValue(v Value) *Credential {
	c.Value = v
	return c
}
func New() *Credential {
	return &Credential{}
}

type CredentialSource interface {
	NameSource
	ValueSource
}

type Credentials interface {
	Get(Name) (Value, error)
	Set(Name, Value) error
}

type Map map[Name]Value

func (m *Map) Get(n Name) (Value, error) {
	return (*m)[n], nil
}
func (m *Map) Set(n Name, v Value) error {
	(*m)[n] = v
	return nil
}
func NewMap() *Map {
	return &Map{}
}
