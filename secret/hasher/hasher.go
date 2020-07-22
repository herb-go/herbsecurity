package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Hasher func(string) (string, error)

var Md5Hasher = func(s string) (string, error) {
	result := md5.Sum([]byte(s))
	return hex.EncodeToString(result[:]), nil
}
var Sha1Hasher = func(s string) (string, error) {
	result := sha1.Sum([]byte(s))
	return hex.EncodeToString(result[:]), nil
}
var Sha256Hasher = func(s string) (string, error) {
	result := sha256.Sum256([]byte(s))
	return hex.EncodeToString(result[:]), nil
}

func GetHasher(name string) (Hasher, error) {
	switch name {
	case "md5":
		return Md5Hasher, nil
	case "sha1":
		return Sha1Hasher, nil
	case "sha256":
		return Sha256Hasher, nil
	}
	return nil, fmt.Errorf("%w (%s)", ErrHasherNotFound, name)
}
