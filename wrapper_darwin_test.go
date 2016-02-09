package service

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

const go_spfc_test_plist string = "<?xml version=\"1.0\" encoding=\"UTF-8\"?> " +
	"<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\"> " +
	"<plist version=\"1.0\"> " +
	"<dict> " +
	"  <key>Label</key> " +
	"  <string>go-spfc-test</string> " +
	"  <key>ProgramArguments</key> " +
	"  <array> " +
	"    <string>bash</string> " +
	"    <string>-c</string> " +
	"    <string>sleep 60</string> " +
	"  </array> " +
	"  <key>UserName</key> " +
	"  <string>'$USER'</string> " +
	"</dict> " +
	"</plist> "

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createService() {
	s := []string{os.Getenv("HOME"), "/Library/LaunchAgents/", servNameTest, ".plist"}
	servNameTestPList := strings.Join(s, "")
	f, err := os.Create(servNameTestPList)
	w := bufio.NewWriter(f)
	_, err = w.WriteString(go_spfc_test_plist)
	w.Flush()
	check(err)
	execLaunchctl("load", servNameTestPList)
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
