// Package service provides Start, Status and Stop functions
package service

import (
	"errors"
	"os/exec"
	"time"
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

	// Handler represents a service instance for execution operation
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

// NewHandler constructs a handler with a given name.
// In linux with sudo true and Mac sudo false
func NewHandler(serviceName string) *Handler {
	return &Handler{sudoDefault(), serviceName}
}

// Start starts a service
func (h *Handler) Start() (Status, error) {
	return h.execService(srvControl.startCmd(h.ServiceName))
}

// StartAndWait starts a service and wait it starts
func (h *Handler) StartAndWait(timeout time.Duration) (Status, error) {
	_, err := h.Start()
	if err != nil {
		return Status{}, err
	}
	return h.waitTimeout(false, timeout)
}

// GetStatus show the status for a service
func (h *Handler) GetStatus() (Status, error) {
	out, err := h.execServiceCmd(srvControl.statusCmd(h.ServiceName))
	return srvControl.parseStatus(out, err)
}

// Stop stops a service
func (h *Handler) Stop() (Status, error) {
	return h.execService(srvControl.stopCmd(h.ServiceName))
}

// StopAndWait stops a service and wait it stops
func (h *Handler) StopAndWait(timeout time.Duration) (Status, error) {
	_, err := h.Stop()
	if err != nil {
		return Status{}, err
	}
	return h.waitTimeout(true, timeout)
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

func (h *Handler) waitTimeout(running bool, timeout time.Duration) (Status, error) {
	timeoutChannel := make(chan Status, 1)
	go func() {
		st, _ := h.GetStatus()
		for getRunningCondition(running, st) {
			st, _ = h.GetStatus()
			time.Sleep(100 * time.Millisecond)
		}
		timeoutChannel <- st
	}()

	select {
	case res := <-timeoutChannel:
		return res, nil
	case <-time.After(timeout):
		return Status{}, errors.New("timeout after " + timeout.String())
	}
}

func getRunningCondition(running bool, st Status) bool {
	if running {
		return st.Running
	}
	return !st.Running
}

func execCmd(cmd string, arg ...string) (string, error) {
	out, err := exec.Command(cmd, arg...).CombinedOutput()
	return string(out), err
}
