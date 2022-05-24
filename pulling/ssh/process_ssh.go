package SSH

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
)

func process_ssh(sshClient *ssh.Client, host string) {

	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	psaux, err := session.Output("ps aux")
	if err != nil {
		panic(err)
	}
	psauxString := string(psaux)

	myStringArray := strings.Split(psauxString, "\n")

	var processList []map[string]string
	flag := 1
	for _, v := range myStringArray {
		if flag == 1 {
			flag = 0
			continue
		}
		smitEachWorld := strings.SplitN(standardizeSpaces(v), " ", 11)
		if len(smitEachWorld) <= 10 {
			break
		}

		temp1 := map[string]string{
			"host":                   host,
			"process.user":           smitEachWorld[0],
			"process.pid":            smitEachWorld[1],
			"process.cpu.percentage": smitEachWorld[2],
			"process.mem.percentage": smitEachWorld[3],
			"process.cmd":            smitEachWorld[10],
		}
		processList = append(processList, temp1)

	}
	result := map[string]interface{}{
		"process": processList,
	}
	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
