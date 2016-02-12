package service

import "fmt"

var fileName = fmt.Sprintf("/etc/init/%s.conf", servNameTest)

const goSpfcTestSystemd string = "[Unit] " +
	"Description=go-spfc-test " +
	" " +
	"[Service] " +
	"TimeoutStartSec=0 " +
	"ExecStart=/bin/sh -c 'while true; do echo Hello World; sleep 1; done' " +
	" " +
	"[Install] " +
	"WantedBy=multi-user.target "

func createService() {
	fmt.Println(fileName)
	cmd := fmt.Sprintf("echo \"script\n  while : ; do sleep 1 ; done\nend script\" > %s", fileName)
	execCmd("sudo", "sh", "-c", cmd)
}

func removeService() {
	execCmd("sudo", "rm", fileName)
}
