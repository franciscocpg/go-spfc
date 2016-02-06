// Package service provides Start, Status and Stop functions
package service

import (
	"errors"
	"os/exec"
)

// Status represents a service status
type Status struct {
	Running bool
	PID     int
}

// Start starts service s
func Start(s string) (Status, error) {
	return execService("start", s)
}

// GetStatus show the status for a given service name (s)
func GetStatus(s string) (Status, error) {
	return status(s)
}

// Stop stops service s
func Stop(s string) (Status, error) {
	return execService("stop", s)
}

func execService(cmd string, s string) (Status, error) {
	out, err := callService(cmd, s)
	var sr Status
	if err != nil {
		return sr, errors.New(out)
	}
	return status(s)
}

func execCmd(cmd string, arg ...string) (string, error) {
	out, err := exec.Command(cmd, arg...).CombinedOutput()
	return string(out), err
}
