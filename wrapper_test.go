package service

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const servNameTest = "go-spfc-test"

var (
	someServiceThatNotExists *Handler
	someServiceThatExists    *Handler
	sudo                     bool
	timeout                  time.Duration
	forceTimeout             time.Duration
)

func init() {
	timeout = 5 * time.Second
	forceTimeout = 1 * time.Millisecond
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

func TestNewHandler(t *testing.T) {
	h := NewHandler("")
	assert.NotNil(t, h)
}

func TestWhenStartAServiceThatDoesNotExist_ShouldGiveAnError(t *testing.T) {
	ss, err := someServiceThatNotExists.Start()
	assert.NotNil(t, err, "Should give an error")
	assert.False(t, ss.Running)
	assert.Equal(t, 0, ss.PID)
}

func TestWhenStartAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	st, err := someServiceThatExists.Start()
	if err != nil {
		panic(err.Error())
	}
	assert.True(t, st.Running)
	assert.NotEqual(t, 0, st.PID)
}

func TestWhenGetStatusForAServiceThatDoesNotExist_ShouldGiveAnError(t *testing.T) {
	st, err := someServiceThatNotExists.GetStatus()
	assert.NotNil(t, err, "Should give an error")
	assert.False(t, st.Running)
	assert.Equal(t, 0, st.PID)
}

func TestWhenGetStatusForAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	st, err := someServiceThatExists.GetStatus()
	if err != nil {
		panic(err.Error())
	}
	assert.True(t, st.Running)
	assert.NotEqual(t, 0, st.PID)
}

func TestWhenStopAServiceThatDoesNotExist_ShouldGiveAnError(t *testing.T) {
	st, err := someServiceThatNotExists.Stop()
	assert.NotNil(t, err, "Should give an error")
	assert.False(t, st.Running)
	assert.Equal(t, 0, st.PID)
}

func TestWhenStopAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	defer removeService()
	st, _ := someServiceThatExists.GetStatus()
	assert.True(t, st.Running)
	assert.NotEqual(t, 0, st.PID)
	st, err := someServiceThatExists.Stop()
	if err != nil {
		panic(err.Error())
	}
	waitStop()
	assert.False(t, st.Running)
	assert.Equal(t, 0, st.PID)
}

func TestWhenStartAndWaitAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	defer removeService()
	createService()
	st, err := someServiceThatExists.StartAndWait(timeout)
	if err != nil {
		panic(err.Error())
	}
	assert.True(t, st.Running)
	assert.NotEqual(t, 0, st.PID)
}

func TestWhenStartAndWaitAServiceThatDoesExistButTimeout_ShouldGiveAnError(t *testing.T) {
	createService()
	_, err := someServiceThatExists.StartAndWait(forceTimeout)
	assert.NotNil(t, err, "Should give an error")
	assert.EqualError(t, err, "timeout after 1ms")
}

func TestWhenStopAndWaitAServiceThatDoesExist_ShoudWorkFine(t *testing.T) {
	defer removeService()
	st, err := someServiceThatExists.StopAndWait(timeout)
	if err != nil {
		panic(err.Error())
	}
	assert.False(t, st.Running)
	assert.Equal(t, 0, st.PID)
}

func TestWhenStopAndWaitAServiceThatDoesExistButTimeout_ShouldGiveAnError(t *testing.T) {
	defer removeService()
	createService()
	st, err := someServiceThatExists.StartAndWait(timeout)
	assert.True(t, st.Running)
	assert.NotEqual(t, 0, st.PID)
	if err != nil {
		panic(err.Error())
	}
	_, err = someServiceThatExists.StopAndWait(forceTimeout)
	assert.NotNil(t, err, "Should give an error")
	assert.EqualError(t, err, "timeout after 1ms")
}
