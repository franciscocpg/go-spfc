// Package service provides Start, Status and Stop functions
package service

import "os/exec"

type StatusResponse struct {
	Running bool
	PID     int
}

// Start starts service s
func Start(s string) (StatusResponse, error) {
	return start(s)
}

// Status show the status for a given service name (s)
func Status(s string) (StatusResponse, error) {
	return status(s)
}

// Stop stops service s
func Stop(s string) (StatusResponse, error) {
	return start(s)
}

func execCmd(cmd string, arg ...string) string {
	out, err := exec.Command(cmd, arg...).CombinedOutput()
	if err != nil {
		panic(string(out))
	}
	return string(out)
}
