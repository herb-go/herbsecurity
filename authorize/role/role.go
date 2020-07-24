package role

type Role struct {
	Name       string
	Attributes *Attributes
}

func (r *Role) Contains(target *Role) bool {

	if r.Name != target.Name {
		return false
	}
	return r.Attributes.Contains(target.Attributes)
}

func (r *Role) WithNewAttribute(name string, value []byte) *Role {
	r.Attributes.Add(name, value)
	return r
}

func (r *Role) WithNewAttributeValues(name string, values ...[]byte) *Role {
	r.Attributes.AddValues(name, values...)
	return r
}

//NewRole create new role by given name.
func NewRole(name string) *Role {
	return &Role{
		Name:       name,
		Attributes: NewAttributes(),
	}
}

type Roles []*Role

func (r *Roles) Data() []*Role {
	return []*Role(*r)
}
func (r *Roles) Append(roles ...*Role) {
	*r = append(*r, roles...)
}

func (r *Roles) FindAll(name string) *Roles {
	result := NewRoles()
	for _, v := range *r {
		if v.Name == name {
			result.Append(v)
		}
	}
	return result
}
func (r *Roles) Contains(target *Roles) bool {
Found:
	for _, v := range *target {
		for _, v2 := range *r {
			if v2.Contains(v) {
				continue Found
			}
		}
		return false
	}
	return true
}

func (r *Roles) Authorize(toauth *Roles) (bool, error) {
	return toauth.Contains(r), nil
}

func (r *Roles) LoadRoles() (*Roles, error) {
	return r, nil
}
func (r *Roles) Clone() *Roles {
	result := make(Roles, len(*r))
	copy(result, *r)
	return &result
}
func NewRoles(r ...*Role) *Roles {
	roles := Roles(r)
	return (&roles).Clone()
}

type RolesLoader interface {
	LoadRoles() (*Roles, error)
}
