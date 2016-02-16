package service

import (
	"errors"
	"strconv"
	"strings"
)

type launchctl struct {
}

const launchctlExec = "launchctl"

func (s launchctl) startCmd(sName string) []string {
	return []string{launchctlExec, "start", sName}
}

func (s launchctl) statusCmd(sName string) []string {
	return []string{launchctlExec, "list", sName}
}

func (s launchctl) stopCmd(sName string) []string {
	return []string{launchctlExec, "stop", sName}
}

func (s launchctl) parseStatus(sData string, err error) (Status, error) {
	var st Status
	if err != nil {
		err = errors.New(string(sData))
	} else {
		lines := strings.Split(string(sData), "\n")
		for _, line := range lines {
			line = strings.Trim(line, "\t")
			if strings.HasPrefix(line, "\"PID\"") {
				st.Running = true
				st.PID, _ = strconv.Atoi(line[8 : len(line)-1])
				break
			}
		}
	}
	return st, err
}
