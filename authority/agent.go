package authority

type Agent string

func (a Agent) AgentData() (Agent, error) {
	return a, nil
}
func (a Agent) String() string {
	return string(a)
}

type AgentSource interface {
	AgentData() (Agent, error)
}
