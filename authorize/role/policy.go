package role

type Policy interface {
	Authorize(*Roles) (bool, error)
}

func Authorize(r *Roles, p ...Policy) (bool, error) {
	if len(p) == 0 {
		return false, nil
	}
	for _, v := range p {
		ok, err := v.Authorize(r)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}
