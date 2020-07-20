package authority

type Payload struct {
	Type string
	Data []byte
}

type Payloads []*Payload
