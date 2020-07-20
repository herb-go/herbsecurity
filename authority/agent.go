package authority

type Agent string

func (a Agent) AgentData() (Agent, error) {
	return a, nil
}

type AgentSource interface {
	AgentData() (Agent, error)
}
