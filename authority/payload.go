package authority

type Payload struct {
	Name string
	Data []byte
}

type Payloads []*Payload

func (p *Payloads) Clone() *Payloads {
	payloads := make(Payloads, len(*p))
	copy(payloads, *p)
	return &payloads
}
func (p *Payloads) Set(name string, data []byte) {
	for k := range *p {
		if (*p)[k].Name == name {
			(*p)[k].Data = data
		}
	}
	*p = append(*p, &Payload{Name: name, Data: data})
}
func (p *Payloads) Add(pl *Payload) {
	p.Set(pl.Name, pl.Data)
}
func (p *Payloads) Load(name string) []byte {
	if p == nil {
		return nil
	}
	for k := range *p {
		if (*p)[k].Name == name {
			return (*p)[k].Data
		}
	}
	return nil
}
func (p *Payloads) LoadString(name string) string {
	return string(p.Load(name))
}

func (p *Payloads) HumanReadabe() HumanReadablePayloads {
	result := HumanReadablePayloads{}
	if p == nil {
		return result
	}
	for _, v := range *p {
		result[v.Name] = string(v.Data)
	}
	return result
}
func NewPayloads() *Payloads {
	return &Payloads{}
}

var PayloadSignSecret = "signsecret"
var PayloadRoles = "roles"

type HumanReadablePayloads map[string]string

func (p HumanReadablePayloads) Payloads() *Payloads {
	result := NewPayloads()
	for k, v := range p {
		result.Set(k, []byte(v))
	}
	return result
}
