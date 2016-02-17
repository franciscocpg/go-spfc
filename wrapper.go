// Package service provides Start, Status and Stop functions
package service

import (
	"errors"
	"os/exec"
)

var (
	srvControl  control
	controlType ControlType
)

type (
	// ControlType represents the service control type
	ControlType int

	// Control is a interface that represents services control (launchctl, initctl, systemctl, etc)
	control interface {
		startCmd(sName string) []string
		stopCmd(sName string) []string
		statusCmd(sName string) []string
		parseStatus(sData string, err error) (Status, error)
	}

	// Execution represents a service instance for execution operation
	Execution struct {
		Sudo        bool
		ServiceName string
	}
	// Status represents a service status
	Status struct {
		Running bool
		PID     int
	}
)

const (
	none ControlType = 1 + iota
	// LaunchCtl - Mac OS implementation (https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/launchctl.1.html)
	LaunchCtl
	// Upstart implementation (http://upstart.ubuntu.com/)
	Upstart
	// SystemD is systemd implementation (https://fedoraproject.org/wiki/Systemd, https://github.com/systemd/systemd)
	SystemD
)

func init() {
	controlType, srvControl = getControlType()
}

// NewExecution constructs a execution with a given name.
// In linux with sudo true and Mac sudo false
func NewExecution(serviceName string) *Execution {
	return &Execution{sudoDefault(), serviceName}
}

// Start starts service
func (e *Execution) Start() (Status, error) {
	return e.execService(srvControl.startCmd(e.ServiceName))
}

// GetStatus show the status for a given service name
func (e *Execution) GetStatus() (Status, error) {
	out, err := e.execServiceCmd(srvControl.statusCmd(e.ServiceName))
	return srvControl.parseStatus(out, err)
}

// Stop stops service
func (e *Execution) Stop() (Status, error) {
	return e.execService(srvControl.stopCmd(e.ServiceName))
}

func (e *Execution) execService(cmdArr []string) (Status, error) {
	out, err := e.execServiceCmd(cmdArr)
	if err != nil {
		return Status{}, errors.New(out)
	}
	return e.GetStatus()
}

func (e *Execution) execServiceCmd(cmdArr []string) (string, error) {
	if e.Sudo {
		return execCmd("sudo", cmdArr...)
	}
	return execCmd(cmdArr[0], cmdArr[1:len(cmdArr)]...)
}

func execCmd(cmd string, arg ...string) (string, error) {
	out, err := exec.Command(cmd, arg...).CombinedOutput()
	return string(out), err
}
