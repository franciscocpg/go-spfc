package service

import (
	"errors"
	"strconv"
	"strings"
)

const initctlExec = "initctl"

type initctl struct {
}

func (s initctl) startCmd(sName string) []string {
	return []string{initctlExec, "start", sName}
}

func (s initctl) statusCmd(sName string) []string {
	return []string{initctlExec, "status", sName}
}

func (s initctl) stopCmd(sName string) []string {
	return []string{initctlExec, "stop", sName}
}

func (s initctl) parseStatus(sData string, err error) (Status, error) {
	var st Status
	lines := strings.Split(sData, "\n")
	if len(lines) == 2 {
		if err != nil {
			err = errors.New(string(sData))
		} else {
			words := strings.Split(string(sData), " ")
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
		}
	}
	return st, nil
}
