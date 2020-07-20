package authority

import (
	"time"
)

type Expiration time.Time

func (e *Expiration) Expired() bool {
	if e == ExpirationNever {
		return false
	}
	return time.Time(*e).After(time.Now())
}

func (e *Expiration) ExpirationData() (*Expiration, error) {
	return e, nil
}

var ExpirationNever *Expiration = nil

type ExpirationSource interface {
	ExpirationData() (*Expiration, error)
}
