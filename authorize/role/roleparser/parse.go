package roleparser

import (
	"strings"
	"sync"

	"github.com/herb-go/herbsecurity/authorize/role"
)

func StringifyRoles(roles *role.Roles) string {
	var result = []string{}
	for k := range *roles {
		result = append(result, stringifyRole((*roles)[k]))
	}
	return strings.Join(result, TokenRoleSep)
}

func stringifyRole(role *role.Role) string {
	rolename := Escape(role.Name)
	attributes := stringifyAttributes(role.Attributes)
	if attributes == "" && role.Name != "" {
		return rolename
	}
	return rolename + TokenAttributesStart + attributes
}

func stringifyAttributes(attributes *role.Attributes) string {
	var attributelist = []string{}
	for _, v := range *attributes {
		attributelist = append(attributelist, stringifyAttribute(v))
	}
	return strings.Join(attributelist, TokenAttributeSep)
}

func stringifyAttribute(a *role.Attribute) string {
	return Escape(a.Keyword) + TokenAttributeVauleStart + Escape(string(a.Value))
}
func parseAttribute(str string) (*role.Attribute, error) {
	var err error
	attr := role.NewAttribute()
	var result = strings.Split(str, TokenAttributeVauleStart)
	switch len(result) {
	case 2:
		attr.Keyword, err = Unescape(result[0])
		if err == nil {
			var value string
			value, err = Unescape(result[1])
			if err == nil {
				attr.Value = []byte(value)
				return attr, nil
			}
		}
	default:
		err = ErrInvalidRole(str)
	}
	return nil, err
}
func parseAttributes(str string) (*role.Attributes, error) {
	a := role.NewAttributes()
	if str == "" {
		return a, nil
	}
	var result = strings.Split(str, TokenAttributeSep)
	for _, v := range result {
		attr, err := parseAttribute(v)
		if err != nil {
			return nil, err
		}
		a.Append(attr)
	}
	return a, nil
}
func parseRole(str string) (*role.Role, error) {
	var err error
	r := role.NewRole("")
	var result = strings.Split(str, TokenAttributesStart)
	switch len(result) {
	case 1:
		r.Name, err = Unescape(result[0])
		if err == nil {
			return r, nil
		}
	case 2:
		r.Name, err = Unescape(result[0])
		if err == nil {
			var attributes *role.Attributes
			attributes, err = parseAttributes(result[1])
			if err == nil {
				r.Attributes = attributes
				return r, nil
			}
		}
	default:
		err = ErrInvalidRole(str)
	}
	return nil, err
}
func ParseRoles(str string) (*role.Roles, error) {
	roles := &role.Roles{}
	rolelist := strings.Split(str, TokenRoleSep)
	for _, v := range rolelist {
		if v == "" {
			continue
		}
		role, err := parseRole(v)
		if err != nil {
			return nil, err
		}
		roles.Append(role)
	}
	return roles, nil
}

var cache sync.Map

var DisableCache bool

func Parse(str string) (*role.Roles, error) {
	if !DisableCache {
		var ok bool
		var r interface{}
		r, ok = cache.Load(str)
		if ok {
			return r.(*role.Roles).Clone(), nil
		}
	}
	roles, err := ParseRoles(str)
	if err != nil {
		return nil, err
	}
	if !DisableCache {
		cache.Store(str, roles.Clone())
	}
	return roles, err
}
