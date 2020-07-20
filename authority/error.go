package authority

import "errors"

var ErrNotFound = errors.New("not found")

var ErrAuthorityDuplicated = errors.New("authority duplicated")

var ErrEmptyAuthority = errors.New("empty authority")

var ErrEmptyPrincipal = errors.New("empty principal")
