package credential

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"
)

var testAuthenticator = AuthenticatorFunc(
	func(c Credentials) (*authority.Auth, error) {
		auth, err := c.Get(NameAuthority)
		if err != nil {
			return nil, err
		}
		passphrase, err := c.Get(NamePassphrase)
		if err != nil {
			return nil, err
		}
		if string(auth) == "testappid" && string(passphrase) == "testtoken" {
			return authority.NewAuth("testappid"), nil
		}
		return nil, nil
	},
	NameAuthority,
	NamePassphrase,
)

func TestVerifier(t *testing.T) {
	appid := New().WithName(NameAuthority).WithValue([]byte("testappid"))
	token := New().WithName(NamePassphrase).WithValue([]byte("testtoken"))
	auth, err := AuthenticateWithDependences(testAuthenticator, appid)
	if auth.Authenticated() || err != nil {
		t.Fatal(auth, err)
	}
	auth, err = AuthenticateWithDependences(testAuthenticator, appid, token)
	if auth.Principal() != "testappid" || err != nil {
		t.Fatal(auth, err)
	}
	auth, err = AuthenticateWithDependences(ForbiddenAuthenticator, appid, token)
	if auth.Authenticated() || err != nil {
		t.Fatal(auth, err)
	}

}
