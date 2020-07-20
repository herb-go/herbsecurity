package keyencoding

import (
	"encoding/base64"

	"github.com/herb-go/herbsecurity/secret"
)

type Encoding struct {
	Encode func(secret.Secret) (string, error)
	Decode func(string) (secret.Secret, error)
}

var NopEncoding = &Encoding{
	Encode: func(s secret.Secret) (string, error) {
		return string(s), nil
	},
	Decode: func(s string) (secret.Secret, error) {
		return []byte(s), nil
	},
}

var Base64Encoding = &Encoding{
	Encode: func(s secret.Secret) (string, error) {
		return base64.StdEncoding.EncodeToString(s), nil
	},
	Decode: func(s string) (secret.Secret, error) {
		return base64.StdEncoding.DecodeString(s)
	},
}
