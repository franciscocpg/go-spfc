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

func sudoDefault() bool {
	return true
}
