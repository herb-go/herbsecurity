package urlencodesign

import (
	"sort"

	"github.com/herb-go/herbsecurity/secret"
	"github.com/herb-go/herbsecurity/secret/hasher"
)

func Sign(h hasher.Hasher, secret secret.Secret, secretfield string, params *Params, asc bool) (string, error) {
	if secretfield == "" {
		return "", ErrSecretFieldEmpty
	}
	p := params.Clone()
	p.Append(secretfield, string(secret))
	if asc {
		sort.Sort(p)
	} else {
		sort.Sort(sort.Reverse(p))
	}
	return h(p.Encode())
}
