package service

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func (e *Execution) status() (Status, error) {
	var st Status
	var out []byte
	var err error
	if e.sudo {
		out, err = exec.Command("sudo", "service", e.ServiceName, "status").CombinedOutput()
	} else {
		out, err = exec.Command("service", e.ServiceName, "status").CombinedOutput()
	}
	if err != nil {
		err = errors.New(string(out))
	} else {
		lines := strings.Split(string(out), " ")
		if len(lines) == 4 {
			if strings.HasPrefix(lines[1], "start/running") {
				st.Running = true
				pid := lines[3][0 : len(lines[3])-1]
				st.PID, err = strconv.Atoi(pid)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return st, err
}

func (e *Execution) callService(cmd string) (string, error) {
	if e.sudo {
		return execCmd("sudo", "service", e.ServiceName, cmd)
	}
	return execCmd("service", e.ServiceName, cmd)
}

func sudoDefault() bool {
	return true
}
