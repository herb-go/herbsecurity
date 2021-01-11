package roleparser

import "testing"

func TestReserved(t *testing.T) {
	var Reserved = []string{
		TokenRoleSep,
		TokenAttributesStart,
		TokenAttributeSep,
		TokenAttributeVauleStart,
	}
	for _, v := range Reserved {
		if Escape(v) == v {
			t.Fatal(v)
		}
	}
}
