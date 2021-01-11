package roleparser

import (
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	err := ErrInvalidRole("teststr")
	if !strings.Contains(err.Error(), "teststr") {
		t.Fatal(err)
	}
}
