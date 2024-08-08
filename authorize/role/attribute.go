package role

import (
	"bytes"
)

type Attribute struct {
	Keyword string
	Value   []byte
}

func (a *Attribute) Equal(b *Attribute) bool {
	if a.Keyword == b.Keyword {
		return bytes.Equal(a.Value, b.Value)
	}
	return false
}
func NewAttribute() *Attribute {
	return &Attribute{}
}

type Attributes []*Attribute

func (a *Attributes) Data() []*Attribute {
	return []*Attribute(*a)
}
func (a *Attributes) Append(attrs ...*Attribute) {
	*a = append(*a, attrs...)
}

func (a *Attributes) Add(keyword string, value []byte) {
	attr := NewAttribute()
	attr.Keyword = keyword
	attr.Value = value
	a.Append(attr)
}
func (a *Attributes) AddValues(keyword string, values ...[]byte) {
	for k := range values {
		a.Add(keyword, values[k])
	}
}
func (a *Attributes) Get(keyword string) []byte {
	length := len(*a)
	if length == 0 {
		return nil
	}
	for i := length - 1; i >= 0; i-- {
		if (*a)[i].Keyword == keyword {
			return (*a)[i].Value
		}
	}
	return nil
}

func (a *Attributes) GetAll(keyword string) *Attributes {
	attrs := NewAttributes()
	for _, v := range *a {
		if v.Keyword == keyword {
			attrs.Append(v)
		}
	}
	return attrs
}

func (a *Attributes) Contains(target *Attributes) bool {
Found:
	for _, v := range *target {
		for _, v2 := range *a {
			if v2.Equal(v) {
				continue Found
			}
		}
		return false
	}
	return true
}
func (a *Attributes) Clone() *Attributes {
	attr := make(Attributes, len(*a))
	copy(attr, *a)
	return &attr
}
func NewAttributes(a ...*Attribute) *Attributes {
	attrs := Attributes(a)
	return attrs.Clone()
}
