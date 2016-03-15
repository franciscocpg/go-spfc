package service

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const servNameTest = "go-spfc-test"

var (
	someServiceThatNotExists *Handler
	someServiceThatExists    *Handler
	sudo                     bool
)

func init() {
	sudoEnv := os.Getenv("GO_SPFC_SUDO_TEST")
	if len(sudoEnv) == 0 {
		sudo = false
	} else {
		sudo, _ = strconv.ParseBool(sudoEnv)
	}
	fmt.Printf("As sudo? %t\n", sudo)
	createService()
	someServiceThatNotExists = &Handler{sudo, "someservicethatnotexists"}
	someServiceThatExists = &Handler{sudo, servNameTest}
}

func TestWhenStartAServiceThatDoesNotExist_ShouldGiveAnError(t *testing.T) {
	ss, err := someServiceThatNotExists.Start()
	assert.NotNil(t, err, "Should give an error")
	assert.False(t, ss.Running)
	assert.Equal(t, 0, ss.PID)
}

func TestWhenStartAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	sr, err := someServiceThatExists.Start()
	if err != nil {
		panic(err.Error())
	}
	assert.True(t, sr.Running)
	assert.NotEqual(t, 0, sr.PID)
}

func TestWhenGetStatusForAServiceThatDoesNotExist_ShouldGiveAnError(t *testing.T) {
	sr, err := someServiceThatNotExists.GetStatus()
	assert.NotNil(t, err, "Should give an error")
	assert.False(t, sr.Running)
	assert.Equal(t, 0, sr.PID)
}

func TestWhenGetStatusForAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	sr, err := someServiceThatExists.GetStatus()
	if err != nil {
		panic(err.Error())
	}
	assert.True(t, sr.Running)
	assert.NotEqual(t, 0, sr.PID)
}

func TestWhenStopAServiceThatDoesNotExist_ShouldGiveAnError(t *testing.T) {
	sr, err := someServiceThatNotExists.Stop()
	assert.NotNil(t, err, "Should give an error")
	assert.False(t, sr.Running)
	assert.Equal(t, 0, sr.PID)
}

func TestWhenStopAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	defer removeService()
	sr, _ := someServiceThatExists.GetStatus()
	assert.True(t, sr.Running)
	assert.NotEqual(t, 0, sr.PID)
	sr, err := someServiceThatExists.Stop()
	if err != nil {
		panic(err.Error())
	}
	assert.False(t, sr.Running)
	assert.Equal(t, 0, sr.PID)
}
