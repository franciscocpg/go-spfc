package service

import "fmt"

var fileName string

const goSpfcTestSystemd string = "\"[Unit]\n " +
	"Description=go-spfc-test\n " +
	"\n " +
	"[Service]\n " +
	"TimeoutStartSec=0\n " +
	"ExecStart=/bin/sh -c 'while true; do sleep 1; done'\n " +
	"\n " +
	"[Install]\n " +
	"WantedBy=multi-user.target\" "

func createService() {
	var text string
	if controlType == SystemD {
		fileName = fmt.Sprintf("/usr/lib/systemd/system/%s.service", servNameTest)
		text = goSpfcTestSystemd
	} else {
		fileName = fmt.Sprintf("/etc/init/%s.conf", servNameTest)
		text = "\"script\n  while : ; do sleep 1 ; done\nend script\""
	}
	fmt.Println(fileName)
	cmd := fmt.Sprintf("echo %s > %s", text, fileName)
	_, err := execCmd("sudo", "sh", "-c", cmd)
	if err != nil {
		panic(err)
	}
}

func removeService() {
	execCmd("sudo", "rm", fileName)
}
