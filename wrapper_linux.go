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
		sOut := string(out)
		lines := strings.Split(sOut, "\n")
		// Upstart
		if len(lines) == 2 {
			words := strings.Split(string(out), " ")
			if len(words) == 4 {
				if strings.HasPrefix(words[1], "start/running") {
					st.Running = true
					pid := words[3][0 : len(words[3])-1]
					st.PID, err = strconv.Atoi(pid)
					if err != nil {
						panic(err)
					}
				}
			}
		} else if strings.HasPrefix(strings.Trim(lines[1], " "), "Loaded") {
			// SystemV
			for _, line := range lines {
				line = strings.Trim(line, " ")
				if strings.HasPrefix(line, "Active") {
					st.Running = strings.Contains(line, "active (running)")
					if !st.Running {
						break
					}
				} else if strings.HasPrefix(line, "Main PID") {
					pid := line[10:len(line)]
					idx := strings.Index(pid, " ")
					pid = pid[0 : idx-1]
					st.PID, _ = strconv.Atoi(pid)
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
