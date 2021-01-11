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
	ok, err = Or().Authorize(nil)
	if ok != true || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = Or(Deny, Deny, Allow).Authorize(nil)
	if ok != true || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = Or(Deny, Deny, Deny).Authorize(nil)
	if ok != false || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = And().Authorize(nil)
	if ok != false || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = And(Deny, Deny, Allow).Authorize(nil)
	if ok != false || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = And(Allow, Allow, Allow).Authorize(nil)
	if ok != true || err != nil {
		t.Fatal(ok, err)
	}
}
