package authority

import "errors"

var ErrAuthorityNotFound = errors.New("authority not found")

var ErrAuthorityDuplicated = errors.New("authority duplicated")

var ErrEmptyAuthority = errors.New("empty authority")

var ErrPrincipalNotFound = errors.New("principal not found")
var ErrEmptyPrincipal = errors.New("empty principal")
var ErrAgentNotFound = errors.New("agent not found")
