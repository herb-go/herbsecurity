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

type PolicyFunc func(*Roles) (bool, error)

func (f PolicyFunc) Authorize(r *Roles) (bool, error) {
	return f(r)
}

func Not(p Policy) Policy {
	return PolicyFunc(func(r *Roles) (bool, error) {
		ok, err := p.Authorize(r)
		if err != nil {
			return false, err
		}
		return !ok, nil
	})
}

func Or(policies ...Policy) Policy {
	return PolicyFunc(func(r *Roles) (bool, error) {
		if len(policies) == 0 {
			return true, nil
		}
		for _, v := range policies {
			ok, err := v.Authorize(r)
			if err != nil {
				return false, err
			}
			if ok {
				return true, nil
			}
		}
		return false, nil
	})
}

func And(policies ...Policy) Policy {
	return PolicyFunc(func(r *Roles) (bool, error) {
		if len(policies) == 0 {
			return false, nil
		}
		for _, v := range policies {
			ok, err := v.Authorize(r)
			if err != nil {
				return false, err
			}
			if !ok {
				return false, nil
			}
		}
		return true, nil
	})
}

var Allow = PolicyFunc(func(r *Roles) (bool, error) {
	return true, nil
})

var Deny = PolicyFunc(func(r *Roles) (bool, error) {
	return false, nil
})
