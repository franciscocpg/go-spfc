package service

func getControlType() (ControlType, control) {
	_, err := execCmd(systemctlExec, "--version")
	if err == nil {
		return SystemD, systemctl{}
	}
	_, err = execCmd(initctlExec, "--version")
	if err == nil {
		return Upstart, initctl{}
	}
	panic("We need systemctl (systemd) or initctl (upstart) implementation")
}

//func (e *Execution) status() (Status, error) {
//	var st Status
//	var out []byte
//	var err error
//	if e.Sudo {
//		out, err = exec.Command("sudo", "service", e.ServiceName, "status").CombinedOutput()
//	} else {
//		out, err = exec.Command("service", e.ServiceName, "status").CombinedOutput()
//	}

//	sOut := string(out)
//	lines := strings.Split(sOut, "\n")
//	// Upstart
//	if len(lines) == 2 {
//		if err != nil {
//			err = errors.New(string(out))
//		} else {
//			words := strings.Split(string(out), " ")
//			if len(words) == 4 {
//				if strings.HasPrefix(words[1], "start/running") {
//					st.Running = true
//					pid := words[3][0 : len(words[3])-1]
//					st.PID, err = strconv.Atoi(pid)
//					if err != nil {
//						panic(err)
//					}
//				}
//			}
//			st.Control = Upstart
//		}
//	} else if strings.HasPrefix(strings.Trim(lines[1], " "), "Loaded") {
//		// systemd
//		if err != nil {
//			errString := err.Error()
//			idx := strings.LastIndex(errString, " ")
//			exitCode, _ := strconv.Atoi(errString[idx+1 : len(errString)])
//			// SystemD status is 3 when service is stopped
//			if exitCode != 3 {
//				return st, err
//			}
//			err = nil
//		}
//		for _, line := range lines {
//			line = strings.Trim(line, " ")
//			if strings.HasPrefix(line, "Active") {
//				st.Running = strings.Contains(line, "active (running)")
//				if !st.Running {
//					break
//				}
//			} else if strings.HasPrefix(line, "Main PID") {
//				pid := line[10:len(line)]
//				idx := strings.Index(pid, " ")
//				pid = pid[0 : idx-1]
//				st.PID, _ = strconv.Atoi(pid)
//			}
//		}
//		st.Control = SystemD
//	}
//	return st, err
//}

//func (e *Execution) callService(cmd string) (string, error) {
//	if e.Sudo {
//		return execCmd("sudo", "service", e.ServiceName, cmd)
//	}
//	return execCmd("service", e.ServiceName, cmd)
//}

func sudoDefault() bool {
	return true
}
