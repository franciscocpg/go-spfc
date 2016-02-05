// Package service provides Start, Status and Stop functions
package service

import (
	"errors"
	"os/exec"
)

type StatusResponse struct {
	Running bool
	PID     int
}

// Start starts service s
func Start(s string) (StatusResponse, error) {
	return execService("start", s)
}

// Status show the status for a given service name (s)
func Status(s string) (StatusResponse, error) {
	return status(s)
}

// Stop stops service s
func Stop(s string) (StatusResponse, error) {
	return execService("stop", s)
}

func execService(cmd string, s string) (StatusResponse, error) {
	out, err := callService(cmd, s)
	var sr StatusResponse
	if err != nil {
		return sr, errors.New(out)
	} else {
		return status(s)
	}
}

func execCmd(cmd string, arg ...string) (string, error) {
	out, err := exec.Command(cmd, arg...).CombinedOutput()
	return string(out), err
}
