package role

import "testing"

func TestPolicy(t *testing.T) {
	rs := NewRoles(NewRole("testrole"), NewRole("testnewrole"))
	policy := Policy(NewRoles(NewRole("testrole")))
	policynewrole := Policy(NewRoles(NewRole("testnewrole")))
	policy2 := Policy(NewRoles(NewRole("testrole2")))
	ok, err := Authorize(rs, policy)
	if err != nil {
		panic(err)
	}
	if !ok {
		t.Fatal(ok)
	}

	ok, err = Authorize(rs, policy2)
	if err != nil {
		panic(err)
	}
	if ok {
		t.Fatal(ok)
	}
	ok, err = Authorize(rs, policy, policynewrole)
	if err != nil {
		panic(err)
	}
	if !ok {
		t.Fatal(ok)
	}
	ok, err = Authorize(rs)
	if err != nil {
		panic(err)
	}
	if ok {
		t.Fatal(ok)
	}
}
