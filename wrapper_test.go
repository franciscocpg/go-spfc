package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const servNameTest = "go-spfc-test"

func init() {
	createService()
}

func TestStatus(t *testing.T) {
	defer removeService()
	sr, err := Status("someservicethatnotexists")
	assert.NotNil(t, err, "Should give an error")
	assert.False(t, sr.Running)
	assert.Equal(t, 0, sr.PID)

	sr, err = Status(servNameTest)
	if err != nil {
		panic(err.Error())
	}
	assert.True(t, sr.Running)
	assert.NotEqual(t, 0, sr.PID)
}
