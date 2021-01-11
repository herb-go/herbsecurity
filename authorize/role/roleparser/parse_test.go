package roleparser

import (
	"testing"

	"github.com/herb-go/herbsecurity/authorize/role"
)

func TestStringify(t *testing.T) {
	var str string
	var roles *role.Roles
	// var err error
	roles = role.NewRoles(
		role.NewRole(""),
		role.NewRole("test1"),
		role.NewRole("test1=1").
			WithNewAttribute("testkey", []byte("testvalue")),
		role.NewRole("test1:1").
			WithNewAttribute("testkey", []byte("testvalue")).
			WithNewAttribute("testkey=2", []byte("testvalue2")),
	)
	str = StringifyRoles(*roles)
	if str != ";test1;test1%3D1:testkey=testvalue;test1%3A1:testkey=testvalue,testkey%3D2=testvalue2" {
		t.Fatal(str)
	}
	rs, err := ParseRoles(str)
	if err != nil {
		t.Fatal(rs, err)
	}
}

func TestCache(t *testing.T) {
	defer func() {
		DisableCache = false
	}()
	var str = ";test1;test1%3D1:testkey=testvalue;test1%3A1:testkey=testvalue,testkey%3D2=testvalue2"
	DisableCache = true
	roles, err := Parse(str)
	if roles == nil || err != nil {
		t.Fatal(roles, err)
	}
	DisableCache = false
	roles, err = Parse(str)
	if roles == nil || err != nil {
		t.Fatal(roles, err)
	}
	roles, err = Parse(str)
	if roles == nil || err != nil {
		t.Fatal(roles, err)
	}
}

func TestInvalid(t *testing.T) {
	roles, err := Parse("1:2:3")
	if roles != nil || err == nil {
		t.Fatal(roles, err)
	}
	roles, err = Parse("a:b=c=d")
	if roles != nil || err == nil {
		t.Fatal(roles, err)
	}
}
