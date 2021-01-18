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

func TestUtil(t *testing.T) {
	var ok bool
	var err error
	ok, err = Allow.Authorize(nil)
	if ok != true || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = Deny.Authorize(nil)
	if ok == true || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = Not(Allow).Authorize(nil)
	if ok == true || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = Any().Authorize(nil)
	if ok != true || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = Any(Deny, Deny, Allow).Authorize(nil)
	if ok != true || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = Any(Deny, Deny, Deny).Authorize(nil)
	if ok != false || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = All().Authorize(nil)
	if ok != false || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = All(Deny, Deny, Allow).Authorize(nil)
	if ok != false || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = All(Allow, Allow, Allow).Authorize(nil)
	if ok != true || err != nil {
		t.Fatal(ok, err)
	}
}
