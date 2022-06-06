package WINRM

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
	"os"
)

func system_winrm(client *winrm.Client, host string) {
	sysUpTime := "(Get-WMIObject win32_operatingsystem).LastBootUpTime;"
	commandForVersion := "(Get-WMIObject win32_operatingsystem).version"
	commandForName := "(Get-WmiObject win32_operatingsystem).name"
	username := "whoami"

	sysUpTime, _, _, errTime := client.RunPSWithString(sysUpTime, "")
	if errTime != nil {
	}
	osVersion, _, _, errVersion := client.RunPSWithString(commandForVersion, "")
	if errVersion != nil {
	}
	osName, _, _, errName := client.RunPSWithString(commandForName, "")
	if errName != nil {
	}
	username, _, _, errUsername := client.RunPSWithString(username, "")
	if errUsername != nil {
	}

	result := map[string]interface{}{
		"host":          host,
		"os.version":    osVersion,
		"os.name":       osName,
		"username":      username,
		"system.uptime": sysUpTime,
		"status.code":   200,
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
