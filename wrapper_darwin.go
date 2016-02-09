package service

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

// Status show the status for a given service name(s)
func (e *Execution) status() (Status, error) {
	out, err := exec.Command("launchctl", "list", e.ServiceName).CombinedOutput()
	var ss Status
	if err != nil {
		err = errors.New(string(out))
	} else {
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			line = strings.Trim(line, "\t")
			if strings.HasPrefix(line, "\"PID\"") {
				ss.Running = true
				ss.PID, _ = strconv.Atoi(line[8 : len(line)-1])
				break
			}
		}
	}
	return ss, err
}

func (e *Execution) callService(cmd string) (string, error) {
	return execCmd("launchctl", cmd, e.ServiceName)
}

func sudoDefault() bool {
	return false
}
