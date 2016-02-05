package service

import "fmt"

var fileName = fmt.Sprintf("/etc/init/%s.conf", servNameTest)

func createService() {
	fmt.Println(fileName)
	cmd := fmt.Sprintf("echo \"script\n  while : ; do sleep 1 ; done\nend script\" > %s", fileName)
	execCmd("sudo", "sh", "-c", cmd)
	Start(servNameTest)
}

func removeService() {
	Stop(servNameTest)
	execCmd("sudo", "rm", fileName)
}
