package urlencodesign

import (
	"testing"

	"github.com/herb-go/herbsecurity/secret"
)

func testHasher(s string) (string, error) {
	return s, nil
}

func TestHash(t *testing.T) {
	p := NewParams()
	p.Append("a", "1")
	p.Append("b", "2")
	sign, err := Sign(testHasher, secret.Secret("secretkey"), "secret", p, true)
	if err != nil {
		panic(err)
	}
	if sign != "a=1&b=2&secret=secretkey" {
		t.Fatal(sign)
	}
	sign, err = Sign(testHasher, secret.Secret("secretkey"), "secret", p, false)
	if err != nil {
		panic(err)
	}
	if sign != "secret=secretkey&b=2&a=1" {
		t.Fatal(sign)
	}
	_, err = Sign(testHasher, secret.Secret("secretkey"), "", p, true)
	if err != ErrSecretFieldEmpty {
		t.Fatal(err)
	}
}
