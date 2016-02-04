package service

import "os/exec"

func createService() {
	execLaunchctl("submit", "-l", servNameTest, "--", "sh", "-c", "while : ; do sleep 1 ; done")
}

func removeService() {
	execLaunchctl("remove", servNameTest)
}

func execLaunchctl(arg ...string) {
	cmd := exec.Command("launchctl", arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(string(out))
	}
}
