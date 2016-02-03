package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintStatus(t *testing.T) {
	_, err := Status("bla")
	assert.NotNil(t, err, "Should give an error")
	sr, err := Status("com.apple.ubd")
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("sr %t %d", sr.Running, sr.PID)
}
