package urlencodesign

import (
	"fmt"
	"sort"

	"github.com/herb-go/herbsecurity/secret"
	"github.com/herb-go/herbsecurity/secret/hasher"
)

var DebugSign bool

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
	encoded := p.Encode()
	result, err := h(encoded)
	if err != nil {
		return "", err
	}
	if DebugSign {
		fmt.Printf("urlencodesign:hashing [%s] as %s\n", encoded, result)
	}
	return result, nil
}
