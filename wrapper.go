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

	// control is a interface that represents services control (launchctl, initctl, systemctl, etc)
	control interface {
		startCmd(sName string) []string
		stopCmd(sName string) []string
		statusCmd(sName string) []string
		parseStatus(sData string, err error) (Status, error)
	}

	// Execution represents a service instance for execution operation
	Handler struct {
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
	// LaunchCtl - Mac OS implementation (https://developer.applh.com/library/mac/documentation/Darwin/Reference/ManPages/man1/launchctl.1.html)
	LaunchCtl
	// Upstart implementation (http://upstart.ubuntu.com/)
	Upstart
	// SystemD is systemd implementation (https://fedoraproject.org/wiki/Systemd, https://github.com/systemd/systemd)
	SystemD
)

func init() {
	controlType, srvControl = getControlType()
}

// NewExecution constructs a execution with a given namh.
// In linux with sudo true and Mac sudo false
func NewHandler(serviceName string) *Handler {
	return &Handler{sudoDefault(), serviceName}
}

// Start starts service
func (h *Handler) Start() (Status, error) {
	return h.execService(srvControl.startCmd(h.ServiceName))
}

// GetStatus show the status for a given service name
func (h *Handler) GetStatus() (Status, error) {
	out, err := h.execServiceCmd(srvControl.statusCmd(h.ServiceName))
	return srvControl.parseStatus(out, err)
}

// Stop stops service
func (h *Handler) Stop() (Status, error) {
	return h.execService(srvControl.stopCmd(h.ServiceName))
}

func (h *Handler) execService(cmdArr []string) (Status, error) {
	out, err := h.execServiceCmd(cmdArr)
	if err != nil {
		return Status{}, errors.New(out)
	}
	return h.GetStatus()
}

func (h *Handler) execServiceCmd(cmdArr []string) (string, error) {
	if h.Sudo {
		return execCmd("sudo", cmdArr...)
	}
	return execCmd(cmdArr[0], cmdArr[1:len(cmdArr)]...)
}

func execCmd(cmd string, arg ...string) (string, error) {
	out, err := exec.Command(cmd, arg...).CombinedOutput()
	return string(out), err
}
