package SSH

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
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
			"process.user":           smitEachWorld[0],
			"process.pid":            smitEachWorld[1],
			"process.cpu.percentage": smitEachWorld[2],
			"process.mem.percentage": smitEachWorld[3],
			"process.cmd":            smitEachWorld[10],
		}
		processList = append(processList, temp1)

	}
	result := map[string]interface{}{
		"host":        host,
		"status":      "success",
		"status.code": 200,
		"process":     processList,
	}
	b, marshalErr := json.Marshal(result)
	if marshalErr != nil {

		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Invalid Polling Json",
			"status.code": 400,
		}

		b, _ := json.Marshal(statusMap)
		encode := base64.StdEncoding.EncodeToString(b)
		fmt.Println(encode)
		os.Exit(0)

	}
	encode := base64.StdEncoding.EncodeToString(b)
	fmt.Println(encode)
}
