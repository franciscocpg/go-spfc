package service

import (
	"strconv"
	"strings"
)

const systemctlExec = "systemctl"

type systemctl struct {
}

func (s systemctl) startCmd(sName string) []string {
	return []string{systemctlExec, "start", sName}
}

func (s systemctl) statusCmd(sName string) []string {
	return []string{systemctlExec, "status", sName}
}

func (s systemctl) stopCmd(sName string) []string {
	return []string{systemctlExec, "stop", sName}
}

func (s systemctl) parseStatus(sData string, err error) (Status, error) {
	var st Status
	lines := strings.Split(sData, "\n")
	if err != nil {
		errString := err.Error()
		idx := strings.LastIndex(errString, " ")
		exitCode, _ := strconv.Atoi(errString[idx+1 : len(errString)])
		// SystemD status is 3 when service is stopped or does not exists
		if exitCode != 3 {
			return st, err
		}
		for _, line := range lines {
			line = strings.Trim(line, " ")
			// When services does not exists then "Loaded" contains "not-found"
			if strings.HasPrefix(line, "Loaded") && strings.Contains(line, "not-found") {
				return st, err
			}
		}
	}
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
	return st, nil
}
