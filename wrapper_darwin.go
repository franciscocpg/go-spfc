package service

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Status show the status for a given service name(s)
func status(s string) (StatusResponse, error) {
	out, err := exec.Command("launchctl", "list", s).CombinedOutput()
	var sr StatusResponse
	if err != nil {
		err = errors.New(string(out))
	} else {
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			line = strings.Trim(line, "\t")
			fmt.Println(line)
			if strings.HasPrefix(line, "\"PID\"") {
				sr.Running = true
				sr.PID, _ = strconv.Atoi(line[8 : len(line)-1])
				break
			}
		}
	}
	return sr, err
}
