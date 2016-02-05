package service

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func status(s string) (StatusResponse, error) {
	var sr StatusResponse
	out, err := exec.Command("sudo", "service", s, "status").CombinedOutput()
	if err != nil {
		err = errors.New(string(out))
	} else {
		lines := strings.Split(string(out), " ")
		if len(lines) == 4 {
			if strings.HasPrefix(lines[1], "start/running") {
				sr.Running = true
				pid := lines[3][0 : len(lines[3])-1]
				sr.PID, err = strconv.Atoi(pid)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return sr, err
}

func callService(cmd string, s string) (string, error) {
	return execCmd("sudo", "service", s, cmd)
}
