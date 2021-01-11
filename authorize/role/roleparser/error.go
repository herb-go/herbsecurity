package roleparser

import (
	"fmt"
)

type ErrInvalidRole string

func (e ErrInvalidRole) Error() string {
	return fmt.Sprintf("roletoken: invalid role [%s]", string(e))
}
