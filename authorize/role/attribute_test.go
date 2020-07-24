package role

import (
	"bytes"
	"testing"
)

func TestAttribute(t *testing.T) {
	attr := NewAttribute()
	attr.Keyword = "testkeyword"
	attr.Value = []byte("testvalue")
	attrnotexists := NewAttribute()
	attrnotexists.Keyword = "testkeyword"
	attrnotexists.Value = []byte("notexists")

	attrs := NewAttributes(attr)
	if attrs.Data()[0] != attr {
		t.Fatal(attrs)
	}
	if !attrs.Contains(NewAttributes(attr)) {
		t.Fatal(attrs)
	}
	if len(attrs.Get("notexists")) != 0 {
		t.Fatal(attrs)
	}

	if attrs.Contains(NewAttributes(attrnotexists)) {
		t.Fatal(attrs)
	}
	if string(attrs.Get("testkeyword")) != "testvalue" {
		t.Fatal(string(attrs.Get("testkeyword")))
	}
	attrs.Add("testkeyword", []byte("newalue"))
	if string(attrs.Get("testkeyword")) != "newalue" {
		t.Fatal(attrs)
	}
	attrsall := attrs.GetAll("testkeyword")
	data := attrsall.Data()
	if len(data) != 2 || bytes.Compare(data[0].Value, attrs.Data()[0].Value) != 0 || bytes.Compare(data[1].Value, attrs.Data()[1].Value) != 0 {
		t.Fatal(data)
	}
	attrs = NewAttributes()
	if len(attrs.Get("testkeyword")) != 0 {
		t.Fatal()
	}
	attrs.AddValues("testkeyword", []byte("value1"), []byte("value2"))
	if len(attrs.GetAll("testkeyword").Data()) != 2 {
		t.Fatal(attrs)
	}
}
