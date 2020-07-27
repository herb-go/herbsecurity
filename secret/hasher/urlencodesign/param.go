package urlencodesign

import (
	"net/url"
	"strings"
)

type Param struct {
	Name  string
	Value string
}

type Params []*Param

func (p *Params) Append(name string, value string) {
	*p = append(*p, &Param{
		Name:  name,
		Value: value,
	})
}

func (p *Params) Len() int {
	return len(*p)
}
func (p *Params) Clone() *Params {
	var params = &Params{}
	*params = make(Params, len(*p))
	copy(*params, *p)
	return params
}

func (p *Params) Encode() string {
	data := []string{}
	for _, v := range *p {
		if v.Value != "" {
			data = append(data, url.QueryEscape(v.Name)+"="+url.QueryEscape(v.Value))
		}
	}
	return strings.Join(data, "&")
}
func NewParams() *Params {
	return &Params{}
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p *Params) Less(i, j int) bool {
	return (*p)[i].Name < (*p)[j].Name
}

// Swap swaps the elements with indexes i and j.
func (p *Params) Swap(i, j int) {
	param := (*p)[i]
	(*p)[i] = (*p)[j]
	(*p)[j] = param
}
