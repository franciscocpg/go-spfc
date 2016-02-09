// Package service provides Start, Status and Stop functions
package service

import (
	"errors"
	"os/exec"
)

type (
	// Execution represents a service instance for execution operation
	Execution struct {
		sudo        bool
		ServiceName string
	}
	// Status represents a service status
	Status struct {
		Running bool
		PID     int
	}
)

// NewExecution constructs a execution with a given name.
// In linux with sudo true and Mac sudo false
func NewExecution(serviceName string) *Execution {
	return &Execution{sudoDefault(), serviceName}
}

// Start starts service s
func (e *Execution) Start() (Status, error) {
	return e.execService("start")
}

// GetStatus show the status for a given service name (s)
func (e *Execution) GetStatus() (Status, error) {
	return e.status()
}

// Stop stops service s
func (e *Execution) Stop() (Status, error) {
	return e.execService("stop")
}

func (e *Execution) execService(cmd string) (Status, error) {
	out, err := e.callService(cmd)
	var sr Status
	if err != nil {
		return sr, errors.New(out)
	}
	return e.status()
}

func execCmd(cmd string, arg ...string) (string, error) {
	out, err := exec.Command(cmd, arg...).CombinedOutput()
	return string(out), err
}
