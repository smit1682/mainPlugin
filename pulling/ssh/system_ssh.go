package SSH

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
)

func system_ssh(sshClient *ssh.Client, host string) {

	session, err := sshClient.NewSession()
	if err != nil {
	}
	sysNameBytes, err := session.Output("hostname")

	sysNameString := string(sysNameBytes)
	sysNameString = strings.TrimRight(sysNameString, "\n")
	session, err = sshClient.NewSession()
	if err != nil {
	}
	osNameBytes, err := session.Output("uname")

	osNameString := string(osNameBytes)
	osNameString = strings.TrimRight(osNameString, "\n")

	session, err = sshClient.NewSession()
	if err != nil {
	}
	osVersionBytes, err := session.Output("hostnamectl | grep 'Operating System'")

	osVersionString := string(osVersionBytes)
	eachw := strings.Split(standardizeSpaces(osVersionString), ":")
	osVersionString = strings.TrimLeft(eachw[1], " ")

	session, err = sshClient.NewSession()
	if err != nil {
	}
	uptimeBytes, err := session.Output("cat /proc/uptime | awk '{print($1)}'")

	uptimeString := string(uptimeBytes)
	uptimeString = strings.TrimRight(uptimeString, "\n")
	session, err = sshClient.NewSession()
	if err != nil {
	}
	threadBytes, err := session.Output("ps -eo nlwp | tail -n +2 | awk '{ num_threads += $1 } END { print num_threads }'")

	threadString := string(threadBytes)
	threadString = strings.TrimRight(threadString, "\n")
	session, err = sshClient.NewSession()
	if err != nil {
	}
	cxcBytes, err := session.Output("vmstat  | awk 'NR==3{print($12)}'")

	cxcString := string(cxcBytes)
	cxcString = strings.TrimRight(cxcString, "\n")
	result := map[string]interface{}{
		"host":                             host,
		"status":                           "success",
		"status.code":                      200,
		"system.name":                      sysNameString,
		"system.os.name":                   osNameString,
		"system.os.version":                osVersionString,
		"system.uptime.seconds":            uptimeString,
		"system.thread.count":              threadString,
		"system.contextSwitch.per.seconds": cxcString,
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
