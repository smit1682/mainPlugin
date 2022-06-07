package SSH

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
)

func memory_ssh(sshClient *ssh.Client, host string) {

	session, err := sshClient.NewSession()
	if err != nil {
	}

	freeMemoryP, err := session.Output("free -b | grep Mem") //free memory

	outputstring := string(freeMemoryP)
	array := strings.Split(standardizeSpaces(outputstring), " ")

	session.Close()

	session, err = sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	usedMemoryP, err := session.Output("free | grep Mem | awk '{ printf(\"%.4f \", $3/$2*100) }'") // used memory
	uString := string(usedMemoryP)
	ufinal := strings.TrimRight(uString, " ")
	session.Close()

	session, err = sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	totalMemoryP, err := session.Output("free | grep Mem | awk '{ printf(\"%.4f \", $4/$2*100) }'") // used memory

	mString := string(totalMemoryP)
	final := strings.TrimRight(mString, " ")
	session.Close()

	result := map[string]interface{}{
		"host":                         host,
		"status":                       "success",
		"status.code":                  200,
		"memory.free.bytes":            array[3],
		"memory.used.bytes":            array[2],
		"memory.available.bytes":       array[6],
		"memory.total.bytes":           array[1],
		"memory.used.percentage.bytes": ufinal,
		"memory.free.percentage.bytes": final,
	}

	marshal, marshalErr := json.Marshal(result)
	if marshalErr != nil {

		statusMap := map[string]interface{}{
			"status":      "error",
			"error":       "Invalid Polling Json",
			"status.code": 400,
		}

		marshal, _ := json.Marshal(statusMap)
		encode := base64.StdEncoding.EncodeToString(marshal)
		fmt.Println(encode)
		os.Exit(0)

	}
	encode := base64.StdEncoding.EncodeToString(marshal)
	fmt.Println(encode)
}
