package role

import "testing"

func TestRole(t *testing.T) {
	role1 := NewRole("testrole")
	rolechina := NewRole("testrole").WithNewAttribute("country", []byte("china"))
	rolechinauas := NewRole("testrole").WithNewAttribute("country", []byte("china")).WithNewAttribute("country", []byte("usa"))
	roleother := NewRole("Othher").WithNewAttribute("country", []byte("china"))
	if role1.Contains(rolechina) {
		t.Fatal()
	}
	if !rolechina.Contains(role1) {
		t.Fatal()
	}
	if rolechina.Contains(roleother) {
		t.Fatal()
	}
	if !rolechinauas.Contains(rolechina) {
		t.Fatal()
	}
}

func TestRoles(t *testing.T) {
	roles := NewRoles(NewRole("testrole"))
	roles.Append(NewRole("testrole2"), NewRole("testrole").WithNewAttribute("country", []byte("china")))
	if len(roles.Data()) != 3 {
		t.Fatal()
	}
	found := roles.FindAll("testrole")
	if len(found.Data()) != 2 {
		t.Fatal()
	}
	if roles.Contains(NewRoles(NewRole("notexists"))) {
		t.Fatal()
	}
	if !roles.Contains(NewRoles(NewRole("testrole").WithNewAttribute("country", []byte("china")))) {
		t.Fatal()
	}
}
