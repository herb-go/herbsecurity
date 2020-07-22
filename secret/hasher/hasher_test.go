package hasher

import (
	"errors"
	"testing"
)

func TestHasher(t *testing.T) {
	var result string
	var err error
	var h Hasher
	data := "12345"
	h, err = GetHasher("md5")
	if err != nil {
		panic(err)
	}
	result, err = h(data)
	if err != nil || result != "827ccb0eea8a706c4c34a16891f84e7b" {
		t.Fatal(result, err)
	}
	h, err = GetHasher("sha1")
	if err != nil {
		panic(err)
	}
	result, err = h(data)
	if err != nil || result != "8cb2237d0679ca88db6464eac60da96345513964" {
		t.Fatal(result, err)
	}
	h, err = GetHasher("sha256")
	if err != nil {
		panic(err)
	}
	result, err = h(data)
	if err != nil || result != "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5" {
		t.Fatal(result, err)
	}
	h, err = GetHasher("notexists")
	if h != nil || errors.Unwrap(err) != ErrHasherNotFound {
		t.Fatal(h, err)
	}

}
