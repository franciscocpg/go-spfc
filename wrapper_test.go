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
	assert.True(t, sr.Running)
	assert.NotEqual(t, 0, sr.PID)
}
