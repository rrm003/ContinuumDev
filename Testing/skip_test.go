package test

import (
	"os"
	"testing"
)

func TestSkip(t *testing.T) {
	if os.Getenv("RRM") == "" {
		t.Skip("... skiping")
	}
	Skip()
}
