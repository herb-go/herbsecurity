package credential

import (
	"github.com/herb-go/herbsecurity/authority"
)

var NamePassphrase = Name("passphrase")

func LoadPassphrase(c Credentials) (authority.Passphrase, error) {
	p, err := c.Get(NamePassphrase)
	if err != nil {
		return "", err
	}
	return authority.Passphrase(p), nil
}

var NameAuthority = Name("authority")

func LoadAuthority(c Credentials) (authority.Authority, error) {
	p, err := c.Get(NameAuthority)
	if err != nil {
		return "", err
	}
	return authority.Authority(p), nil
}

var NameAppID = Name("appid")

func LoadAppID(c Credentials) (authority.Authority, error) {
	p, err := c.Get(NameAppID)
	if err != nil {
		return "", err
	}
	return authority.Authority(p), nil
}

var NameSecret = Name("secret")

func LoadSecret(c Credentials) (authority.Passphrase, error) {
	p, err := c.Get(NameSecret)
	if err != nil {
		return "", err
	}
	return authority.Passphrase(p), nil
}

var NameTimestamp = Name("timestamp")

func LoadTimestamp(c Credentials) (string, error) {
	p, err := c.Get(NameTimestamp)
	if err != nil {
		return "", err
	}
	return string(p), nil
}

var NameSign = Name("sign")

func LoadSign(c Credentials) (string, error) {
	p, err := c.Get(NameSign)
	if err != nil {
		return "", err
	}
	return string(p), nil
}
