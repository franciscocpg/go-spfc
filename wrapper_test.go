package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const servNameTest = "go-spfc-test"

func init() {
	createService()
}

func TestWhenGetStatusForAServiceThatDoesNotExist_ShouldGiveAnError(t *testing.T) {
	sr, err := GetStatus("someservicethatnotexists")
	assert.NotNil(t, err, "Should give an error")
	assert.False(t, sr.Running)
	assert.Equal(t, 0, sr.PID)
}

func TestWhenGetStatusForAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	defer removeService()
	sr, err := GetStatus(servNameTest)
	if err != nil {
		panic(err.Error())
	}
	assert.True(t, sr.Running)
	assert.NotEqual(t, 0, sr.PID)
}
