package authority

type Authentication struct {
	Principal
	Agent
}

func NewAuthentication() *Authentication {
	return &Authentication{}
}
