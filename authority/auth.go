package authority

type Auth struct {
	principal Principal
	agent     Agent
	authority Authority
	payloads  *Payloads
}

func (a *Auth) Principal() Principal {
	if a == nil {
		return ""
	}
	return a.principal
}
func (a *Auth) String() string {
	return string(a.Principal())
}
func (a *Auth) WithAgent(agent Agent) *Auth {
	a.agent = agent
	return a
}
func (a *Auth) Agent() Agent {
	return a.agent
}
func (a *Auth) WithPayloads(payloads *Payloads) *Auth {
	a.payloads = payloads
	return a
}
func (a *Auth) Payloads() *Payloads {
	return a.payloads
}
func (a *Auth) WithAuthority(authority Authority) *Auth {
	a.authority = authority
	return a
}
func (a *Auth) Authority() Authority {
	return a.authority
}

func (a *Auth) Authenticated() bool {
	if a == nil || a.Principal() == "" {
		return false
	}
	return true
}
func NewAuth(p Principal) *Auth {
	return &Auth{
		principal: p,
	}
}

type AuthResult interface {
	Auth() *Auth
}
