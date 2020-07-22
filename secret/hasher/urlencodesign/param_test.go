package urlencodesign

import (
	"sort"
	"testing"
)

func TestParam(t *testing.T) {
	p := NewParams()
	p.Append("b", "2")
	p.Append("a", "1")
	p2 := p.Clone()
	sort.Sort(p2)
	if (*p)[0].Name != "b" || (*p)[1].Name != "a" {
		t.Fatal(p)
	}
	if (*p2)[0].Name != "a" || (*p2)[1].Name != "b" {
		t.Fatal(p2)
	}
	if p.Encode() != "b=2&a=1" {
		t.Fatal(p.Encode())
	}
}
