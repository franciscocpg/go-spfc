package service

func getControlType() (ControlType, control) {
	return LaunchCtl, launchctl{}
}

func sudoDefault() bool {
	return false
}
